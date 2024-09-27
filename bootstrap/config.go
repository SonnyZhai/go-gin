package bootstrap

import (
	"fmt"
	"go-gin/global"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitializeConfig() *viper.Viper {
	// 初始化 viper
	v := viper.New()

	// 优先读取环境变量指定的配置文件路径
	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// 根据APP_ENV环境变量加载相应的配置文件
		env := os.Getenv("APP_ENV")
		switch env {
		case "prod":
			v.SetConfigName("config.prod")
			v.SetConfigType("yaml")
			//. 表示当前目录，也就是项目根目录
			v.AddConfigPath(".")
		case "test":
			v.SetConfigName("config.test")
			v.SetConfigType("yaml")
			v.AddConfigPath(".")
		default:
			v.SetConfigName("settings")
			v.SetConfigType("toml")
			v.AddConfigPath(".")
		}
	}

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %s", err)
	}

	// 监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改并重新加载", in.Name)
		// 重新加载配置文件
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println("重新加载配置文件失败: ", err)
		}
	})

	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		log.Fatalf("配置赋值给全局变量失败: %s", err)
	}

	return v
}
