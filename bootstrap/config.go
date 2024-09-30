package bootstrap

import (
	"go-gin/cons"
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
	configPath := os.Getenv(cons.CONFIG_PATH)
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// 根据APP_ENV环境变量加载相应的配置文件
		env := os.Getenv(cons.APP_ENV)
		switch env {
		case cons.ENV_PROD:
			v.SetConfigName(cons.CONFIG_PROD)
			v.SetConfigType(cons.YAML_TYPE)
			//. 表示当前目录，也就是项目根目录
			v.AddConfigPath(cons.DOT)
		case cons.ENV_TEST:
			v.SetConfigName(cons.CONFIG_TEST)
			v.SetConfigType(cons.YAML_TYPE)
			v.AddConfigPath(cons.DOT)
		default:
			v.SetConfigName(cons.CONFIG_DEV)
			v.SetConfigType(cons.TOML_TYPE)
			v.AddConfigPath(cons.DOT)
		}
	}

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf(cons.FATAL_READ_CONFIG+cons.STRING_PLACEHOLDER, err)
	}

	// 监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println(cons.INFO_MODIFY_CONFIG, in.Name)
		// 重新加载配置文件
		if err := v.Unmarshal(&global.App.Config); err != nil {
			log.Println(cons.ERROR_RELOAD_CONFIG, err)
		}
	})

	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		log.Fatalf(cons.FATAL_CONFIG_TO_GLOBAL+cons.STRING_PLACEHOLDER, err)
	}

	return v
}
