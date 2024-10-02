package routes

import (
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

	// 模拟耗时接口
	router.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})
}
