package bootstrap

import (
	"go-gin/cons"
	"go-gin/global"
	"strconv"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + cons.COLON + strconv.Itoa(global.App.Config.Redis.Port),
		DB:       global.App.Config.Redis.DB,
		Password: global.App.Config.Redis.Password,
	})
	_, err := client.Ping().Result()
	if err != nil {
		global.App.Log.Error(cons.ERROR_REDIS_CONNECTION, zap.Any("err", err))
		return nil
	}
	return client
}
