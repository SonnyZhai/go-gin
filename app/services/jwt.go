package services

import (
	"go-gin/global"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	jwt.StandardClaims
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
func (jwtService *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	// 创建一个新的 token 对象
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTtl, // 过期时间
				Id:        user.GetUid(),                                    // 用户 ID
				Issuer:    GuardName,                                        // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,                         // 生效时间
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
