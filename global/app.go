package global

import (
	"go-gin/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 定义 Application 结构体，用来存放一些项目启动时的变量，便于调用
type Application struct {
	// ConfigViper 用来存放 viper 实例
	ConfigViper *viper.Viper
	// Config 用来存放 config.Configuration 结构体
	Config config.Configuration
	// Log 用来存放 zap.Logger 实例
	Log *zap.Logger
	// DB 用来存放 gorm.DB 实例
	DB *gorm.DB
	// Redis 用来存放 redis.Client 实例
	Redis *redis.Client
	// S3Client 用来存放 s3.Client 实例
	S3 *s3.Client
}

var App = new(Application)
