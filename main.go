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

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	global.App.Log.Info("数据库初始化成功")
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

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
