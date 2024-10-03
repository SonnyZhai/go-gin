package routes

import (
	"go-gin/app/controller/app"
	"go-gin/app/middleware"
	"go-gin/app/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	// 测试接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 模拟耗时接口, 测试优雅关机
	router.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})
}

// 用户分组路由
func SetUserGroupRoutes(router *gin.RouterGroup) {
	router.POST("/register", app.Register)
	router.POST("/login", app.Login)
}

// 需要认证才能访问的接口
func SetAuthGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.GET("/userInfo", app.GetUserInfo)
		authRouter.POST("/logout", app.Logout)
	}
}
