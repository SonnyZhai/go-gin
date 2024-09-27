package global

import (
	"go-gin/config"

	"github.com/spf13/viper"
)

// 定义 Application 结构体，用来存放一些项目启动时的变量，便于调用
type Application struct {
	// ConfigViper 用来存放 viper 实例
	ConfigViper *viper.Viper
	// Config 用来存放 config.Configuration 结构体
	Config config.Configuration
}

var App = new(Application)
