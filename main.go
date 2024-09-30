package main

import (
	"go-gin/bootstrap"
	"go-gin/cons"
	"go-gin/global"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置文件
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info(cons.INFO_LOG_INIT_SUCCESS)

	// 创建一个默认的路由引擎
	r := gin.Default()

	// 测试GET请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 启动服务
	r.Run(cons.COLON + strconv.Itoa(global.App.Config.App.Port))

}
