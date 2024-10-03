// 封装分布式锁
package global

import (
	"context"
	"go-gin/cons"
	"go-gin/utils"
	"time"

	"github.com/go-redis/redis/v8"
)

type Interface interface {
	Get() bool                // 获取锁
	Block(seconds int64) bool // 阻塞获取锁
	Release() bool            // 释放锁
	ForceRelease()            // 强制释放锁
}

type lock struct {
	context context.Context // 上下文
	name    string          // 锁名称
	owner   string          // 锁标识
	seconds int64           // 有效期
}

// 生成锁
func Lock(name string, seconds int64) Interface {
	return &lock{
		context.Background(),
		name,
		utils.RandString(16),
		seconds,
	}
}

// 获取锁
func (l *lock) Get() bool {
	return App.Redis.SetNX(l.context, l.name, l.owner, time.Duration(l.seconds)*time.Second).Val()
}

// 阻塞一段时间，尝试获取锁
func (l *lock) Block(seconds int64) bool {
	starting := time.Now().Unix()
	for {
		if !l.Get() {
			time.Sleep(time.Duration(1) * time.Second)
			if time.Now().Unix()-seconds >= starting {
				return false
			}
		} else {
			return true
		}
	}
}

// 释放锁
func (l *lock) Release() bool {
	luaScript := redis.NewScript(cons.ReleaseLockLuaScript)
	result := luaScript.Run(l.context, App.Redis, []string{l.name}, l.owner).Val().(int64)
	return result != 0
}

// 强制释放锁
func (l *lock) ForceRelease() {
	App.Redis.Del(l.context, l.name).Val()
}
