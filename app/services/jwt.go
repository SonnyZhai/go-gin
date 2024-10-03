package services

import (
	"go-gin/cons"
	"go-gin/global"
	"go-gin/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtService struct {
}

var JwtService = new(jwtService)

// 所有需要颁发 token 的用户模型必须实现这个接口
type JwtUser interface {
	GetUid() string
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.RegisteredClaims
}

// Token 令牌
const (
	TokenType    = "bearer" // 令牌类型
	AppGuardName = "app"    // 守卫名称
)

// TokenOutPut 令牌输出
type TokenOutPut struct {
	AccessToken string `json:"access_token"` // 访问令牌
	ExpiresIn   int    `json:"expires_in"`   // 过期时间
	TokenType   string `json:"token_type "`  // 令牌类型
}

// CreateToken 生成 Token
func (jwtService *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, token *jwt.Token, err error) {
	// 创建一个新的 token 对象
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.App.Config.Jwt.JwtTtl) * time.Second)), // 过期时间
				ID:        user.GetUid(),                                                                                 // 用户 ID
				Issuer:    GuardName,                                                                                     // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: jwt.NewNumericDate(time.Now().Add(-1000 * time.Second)),                                       // 生效时间
			},
		},
	)

	// 生成 token 字符串
	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))

	// 返回 token 数据
	tokenData = TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.JwtTtl),
		TokenType,
	}
	return
}

// 获取黑名单缓存 key
func (jwtService *jwtService) getBlackListKey(tokenStr string) string {
	return cons.JWT_BLACK_LIST + cons.COLON + utils.MD5([]byte(tokenStr))
}

// JoinBlackList token 加入黑名单
func (jwtService *jwtService) JoinBlackList(token *jwt.Token) (err error) {
	// 获取当前的 Unix 时间戳
	nowUnix := time.Now().Unix()
	// 计算 JWT 令牌的剩余有效时间
	expiresAt := token.Claims.(*CustomClaims).ExpiresAt.Unix()
	timer := time.Duration(expiresAt-nowUnix) * time.Second
	//将当前时间作为缓存 value 值，将 token 剩余时间设置为缓存有效期，这样就可以保证 token 在缓存中的有效期和 token 的有效期一致
	err = global.App.Redis.SetNX(jwtService.getBlackListKey(token.Raw), nowUnix, timer).Err()
	return
}

// IsInBlacklist token 是否在黑名单中
func (jwtService *jwtService) IsInBlacklist(tokenStr string) bool {
	// 从 Redis 获取黑名单信息：如果获取到了，说明 token 在黑名单中
	joinUnixStr, err := global.App.Redis.Get(jwtService.getBlackListKey(tokenStr)).Result()
	if err != nil {
		// 如果获取 Redis 值时出错，返回 false
		return false
	}
	// 将从 Redis 获取的字符串值 joinUnixStr 解析为整数时间戳 joinUnix
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true
}
