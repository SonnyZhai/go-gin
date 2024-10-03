package utils

import (
	"math/rand"
	"time"
)

// 随机字符串用于生成锁标识，防止任何客户端都能解锁
func RandString(len int) string {
	// 使用当前时间的纳秒数作为种子创建一个新的随机数源。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 创建一个新的随机数生成器 r
	bytes := make([]byte, len)
	// 生成一个 0 到 25 之间的随机整数，并加上 65。这样生成的整数在 ASCII 码表中对应大写字母 A 到 Z（65 到 90）
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
