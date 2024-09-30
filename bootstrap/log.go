package bootstrap

import (
	"go-gin/cons"
	"go-gin/global"
	"go-gin/utils"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitializeLog() *zap.Logger {
	// 创建根目录
	createRootDir()

	// 设置日志等级
	setLogLevel()

	// 设置日志输出格式
	if global.App.Config.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	// 初始化 zap
	return zap.New(getZapCore(), options...)

}

func createRootDir() {
	// 判断根目录是否存在
	if ok, _ := utils.CheckPathExists("global.App.Config.Log.RootDir"); !ok {
		_ = os.Mkdir(global.App.Config.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.App.Config.Log.Level {
	case cons.LogLevelDebug:
		// 设置日志等级
		level = zap.DebugLevel
		// 设置堆栈跟踪
		options = append(options, zap.AddStacktrace(level))
	case cons.LogLevelInfo:
		level = zap.InfoLevel
	case cons.LogLevelWarn:
		level = zap.WarnLevel
	case cons.LogLevelError:
		level = zap.ErrorLevel
	case cons.LogLevelDPanic:
		level = zap.DPanicLevel
	case cons.LogLevelPanic:
		level = zap.PanicLevel
	case cons.LogLevelFatal:
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format(cons.LEFT_SQUARE_BRACKET + cons.DateTimeFormat + cons.RIGHT_SQUARE_BRACKET))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Config.App.Env.Name + cons.DOT + l.String())
	}

	// 设置编码器
	if global.App.Config.Log.Format == cons.JSON_TYPE {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.App.Config.Log.RootDir + cons.SLASH + global.App.Config.Log.Filename,
		MaxSize:    global.App.Config.Log.MaxSize,
		MaxBackups: global.App.Config.Log.MaxBackups,
		MaxAge:     global.App.Config.Log.MaxAge,
		Compress:   global.App.Config.Log.Compress,
	}

	return zapcore.AddSync(file)
}
