package bootstrap

import (
	"go-gin/cons"
	"go-gin/global"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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

	// 读取本地配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf(cons.FATAL_READ_CONFIG+cons.STRING_PLACEHOLDER, err)
	}

	// 将本地配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		log.Fatalf(cons.FATAL_CONFIG_TO_GLOBAL+cons.STRING_PLACEHOLDER, err)
	}

	// 加载数据库配置
	loadDatabaseConfig(v)

	// 通过 etcd 加载远程配置
	configByEtcd(v)

	// 启动协程监听本地配置文件变化
	go watchLocalConfig(v)

	// 启动协程监听远程配置文件变化
	go watchRemoteConfig(v)

	return v
}

// 根据database字段值加载相应的数据库配置
func loadDatabaseConfig(v *viper.Viper) {
	switch global.App.Config.Database {
	case cons.DATABASE_TYPE_POSTGRESQL:
		if err := v.UnmarshalKey(cons.DATABASE_TYPE_POSTGRESQL, &global.App.Config.PostgresDB); err != nil {
			log.Fatalf(cons.FATAL_DECODE_POSTGRES+cons.STRING_PLACEHOLDER, err)
		}
	case cons.DATABASE_TYPE_MYSQL:
		if err := v.UnmarshalKey(cons.DATABASE_TYPE_MYSQL, &global.App.Config.MysqlDB); err != nil {
			log.Fatalf(cons.FATAL_DECODE_MYSQL+cons.STRING_PLACEHOLDER, err)
		}
	default:
		log.Fatalf(cons.ERROR_DB_TYPE_UNSUPPORT+cons.STRING_PLACEHOLDER, global.App.Config.Database)
	}
}

// 通过 etcd 加载远程配置
func configByEtcd(v *viper.Viper) {
	// etcd 配置
	etcdAddr := os.Getenv(cons.ETCD_ENV_ADDR)

	if etcdAddr == "" {
		log.Fatal(cons.FATAL_ETCD_ADDR_PROVIDER)
	}

	var err error

	// 添加远程配置提供者
	if err = v.AddRemoteProvider(cons.ETCD_VERSION, etcdAddr, cons.ETCD_CONFIG_PATH); err != nil {
		log.Fatal(cons.FATAL_ADD_REMOTE_PROVIDER, err)
	}

	v.SetConfigType(cons.TOML_TYPE)

	// 读取远程配置
	if err = v.ReadRemoteConfig(); err != nil {
		log.Fatal(cons.FATAL_READ_REMOTE_CONFIG, err)
	}

	// 将远程配置赋值给全局变量
	if err = v.UnmarshalKey(cons.OSS_R2_NAME, &global.App.Config.Etcd); err != nil {
		log.Fatal(cons.FATAL_REMOTE_VALUE_TO_CONF, err)
	}
}

// 监听本地配置文件变化
func watchLocalConfig(v *viper.Viper) {
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println(cons.INFO_MODIFY_CONFIG, in.Name)
		// 重新加载配置文件
		if err := v.Unmarshal(&global.App.Config); err != nil {
			log.Println(cons.ERROR_RELOAD_CONFIG, err)
		}
	})
}

// 监听远程配置文件变化
func watchRemoteConfig(v *viper.Viper) {
	for {

		// 每隔30秒监控一次远程配置
		time.Sleep(time.Second * 30)

		// 重新加载远程配置
		err := v.WatchRemoteConfig()
		if err != nil {
			log.Printf(cons.FATAL_READ_REMOTE_CONFIG+cons.STRING_PLACEHOLDER, err)
			continue
		}

		if err := v.UnmarshalKey(cons.OSS_R2_NAME, &global.App.Config.Etcd); err != nil {
			log.Printf(cons.FATAL_REMOTE_VALUE_TO_CONF+cons.STRING_PLACEHOLDER, err)
			continue
		}
	}
}
