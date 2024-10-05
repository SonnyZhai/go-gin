package main

import (
	"go-gin/bootstrap"
	"go-gin/cons"
	"go-gin/global"
)

func main() {
	// 初始化配置文件
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info(cons.INFO_LOG_INIT_SUCCESS)

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	global.App.Log.Info(cons.INFO_DB_CONNECT)
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化 Redis
	global.App.Redis = bootstrap.InitializeRedis()
	global.App.Log.Info(cons.INFO_REDIS_CONNECTION)

	// 初始化 S3
	global.App.S3 = bootstrap.InitializeS3()
	global.App.Log.Info(cons.INFO_S3_CONNECTION)

	// 启动服务
	bootstrap.RunServer()

}
