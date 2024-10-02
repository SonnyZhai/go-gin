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
	global.App.Log.Info("数据库初始化成功")
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 启动服务
	bootstrap.RunServer()

}
