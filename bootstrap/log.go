package bootstrap

import (
	"go-gin/global"
	"go-gin/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitializeLog() *zap.Logger {
	// 创建根目录
	createRootDir()

}

func createRootDir() {
	// 判断根目录是否存在
	if ok, _ := utils.CheckPathExists("global.App.Config.Log.RootDir"); !ok {
		_ = os.Mkdir(global.App.Config.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.App.Config.Log.Level {
	case "debug":
		// 设置日志等级
		level = zap.DebugLevel
		// 设置堆栈跟踪
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}
