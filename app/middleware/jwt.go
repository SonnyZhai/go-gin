package middleware

import (
	"go-gin/app/services"
	"go-gin/cons"
	"go-gin/errors"
	"go-gin/global"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get(cons.API_AUTH_NAME)
		// 如果 token 为空
		if tokenStr == "" {
			errors.HandleErrorWithContext(c, 401, cons.ERROR_CODE_SERVER_USER_INVALID_TOKEN, cons.ERROR_EMPTY_TOKEN, nil, nil)
			c.Abort()
			return
		}

		// tokenStr[len(services.TokenType)+1:] 表示 token 字符串中去掉 "bearer " 后的部分
		tokenStr = tokenStr[len(services.TokenType)+1:]
		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		// Token 校验失败，或者 token 在黑名单中
		if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
			errors.HandleErrorWithContext(c, 401, cons.ERROR_CODE_SERVER_USER_INVALID_TOKEN, cons.ERROR_CLAIMS_TOKEN, nil, nil)
			c.Abort()
			return
		}

		// 从 token 中获取用户信息
		claims := token.Claims.(*services.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			errors.HandleErrorWithContext(c, 401, cons.ERROR_CODE_SERVER_USER_INVALID_TOKEN, cons.ERROR_INVALID_ISSUER, nil, nil)
			c.Abort()
			return
		}

		// token 续签, 如果 token 过期时间小于刷新时间，则重新颁发 token
		if claims.ExpiresAt.Unix()-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
			// 生成一个锁对象，用于确保在续签过程中不会有其他并发请求干扰
			lock := global.Lock(cons.REFRESH_TOKEN_LOCK, global.App.Config.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				// 从 token 中获取用户信息
				user, err := services.JwtService.GetUserInfo(GuardName, claims.ID)
				if err != nil {
					// 如果获取用户信息失败，记录错误日志并释放锁。
					global.App.Log.Error(err.Error())
					lock.Release()
				} else {
					// 生成新的令牌，并将新的令牌信息添加到响应头中。
					tokenData, _, _ := services.JwtService.CreateToken(GuardName, user)
					c.Header(cons.API_REFRESH_TOKEN, tokenData.AccessToken)
					c.Header(cons.API_REFRESH_TOKEN_EXPIRE, strconv.Itoa(tokenData.ExpiresIn))
					_ = services.JwtService.JoinBlackList(token)
				}
			}
		}

		// 存储 token 和用户 ID
		c.Set(cons.API_TOKEN_NAME, token)
		c.Set(cons.API_USER_ID, claims.ID)
	}
}
