package middleware

import (
	"go-gin/app/services"
	"go-gin/cons"
	"go-gin/errors"
	"go-gin/global"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
		if err != nil {
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

		// 存储 token 和用户 ID
		c.Set(cons.API_TOKEN_NAME, token)
		c.Set(cons.API_USER_ID, claims.Id)
	}
}
