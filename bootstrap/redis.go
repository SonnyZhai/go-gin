package bootstrap

import (
	"context"
	"go-gin/cons"
	"go-gin/global"
	"strconv"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + cons.COLON + strconv.Itoa(global.App.Config.Redis.Port),
		DB:       global.App.Config.Redis.DB,
		Password: global.App.Config.Redis.Password,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error(cons.ERROR_REDIS_CONNECTION, zap.Any("err", err))
		return nil
	}
	return client
}
