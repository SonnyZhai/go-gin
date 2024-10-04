package bootstrap

import (
	"bytes"
	"context"
	"fmt"
	"go-gin/cons"
	"go-gin/global"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	clientv3 "go.etcd.io/etcd/client/v3"
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
	// 启动协程监听本地配置文件变化
	go watchLocalConfig(v)

	// 通过 etcd 加载远程配置
	configByEtcd(v)
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

	// 创建 etcd 客户端
	client := newEtcdClient()

	// 使用 etcd 客户端读取配置
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 释放资源
	defer cancel()

	// 读取配置
	resp, err := client.Get(ctx, cons.ETCD_CONFIG_PATH)
	if err != nil {
		log.Fatalf(cons.FATAL_READ_REMOTE_CONFIG+cons.STRING_PLACEHOLDER, err)
	}
	if len(resp.Kvs) == 0 {
		log.Fatal(cons.FATAL_FOUND_NO_CONFIG)
	}

	v.SetConfigType(cons.TOML_TYPE)

	// 读取远程配置
	err = v.ReadConfig(bytes.NewReader(resp.Kvs[0].Value))
	if err != nil {
		log.Fatalf(cons.FATAL_PARSE_REMOTE_CONFIG+cons.STRING_PLACEHOLDER, err)
	}

	// 将远程配置赋值给全局变量
	if err = v.UnmarshalKey(cons.OSS_R2_NAME, &global.App.Config.Etcd); err != nil {
		log.Fatal(cons.FATAL_REMOTE_VALUE_TO_CONF, err)
	}

	fmt.Println(global.App.Config)
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

	// 创建 etcd 客户端
	client := newEtcdClient()
	defer client.Close()

	for {
		log.Println(cons.INFO_START_WATCH_REMOTE_CONFIG)

		watchChan := client.Watch(context.Background(), cons.ETCD_CONFIG_PATH, clientv3.WithPrefix())

		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				log.Printf(cons.INFO_CONFIG_CHANGED+cons.STRING_PLACEHOLDER_N, event.Kv.Key)
				updateConfig(v, client)
			}
		}

		log.Println(cons.INFO_WATCHING_CHANNEL_RESTART)
		// 避免在出错时立即重启
		time.Sleep(time.Second)
	}
}

func updateConfig(v *viper.Viper, client *clientv3.Client) {
	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 获取最新的配置
	resp, err := client.Get(ctx, cons.ETCD_CONFIG_PATH)
	if err != nil {
		log.Printf(cons.FATAL_GET_NEW_REMOTE_CONFIG+cons.STRING_PLACEHOLDER_N, err)
		return
	}

	if len(resp.Kvs) == 0 {
		log.Println(cons.FATAL_FOUND_NO_CONFIG)
		return
	}

	// 使用 Viper 解析新的配置
	if err := v.ReadConfig(bytes.NewReader(resp.Kvs[0].Value)); err != nil {
		log.Printf(cons.FATAL_PARSE_NEW_REMOTE_CONFIG+cons.STRING_PLACEHOLDER_N, err)
		return
	}

	// 更新全局配置
	if err := v.UnmarshalKey(cons.OSS_R2_NAME, &global.App.Config.Etcd); err != nil {
		log.Printf(cons.FATAL_ETCD_CONFIG_TO_GLOBAL+cons.STRING_PLACEHOLDER_N, err)
		return
	}

	log.Println(cons.INFO_UPDATE_CONFIG_SUCCESS)
}

func newEtcdClient() *clientv3.Client {
	// etcd 配置
	etcdAddr := os.Getenv(cons.ETCD_ENV_ADDR)
	etcdPassword := os.Getenv(cons.ETCD_ENV_PASSWORD)

	if etcdAddr == "" {
		log.Fatal(cons.FATAL_ETCD_ADDR_PROVIDER)
	}
	// 创建 etcd 客户端配置
	cfg := clientv3.Config{
		Endpoints:            []string{etcdAddr},
		Username:             cons.ETCD_USERNAME,
		Password:             etcdPassword,
		DialTimeout:          10 * time.Second,
		DialKeepAliveTime:    10 * time.Second, // 设置保活时间，保持连接活跃
		DialKeepAliveTimeout: 30 * time.Second, // 保活超时时间
		PermitWithoutStream:  true,             // 没有流时允许创建客户端
	}

	// 创建 etcd 客户端
	client, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalf(cons.FATAL_CREATE_ETCD_CLIENT+cons.STRING_PLACEHOLDER, err)
	}

	return client
}
