package db

import (
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

// InitMockRedis 回傳一個模擬 Redis 的 client
func InitMockRedis() *redis.Client {
	mr, err := miniredis.Run() // 啟動一個模擬的 Redis 伺服器
	if err != nil {
		log.Fatalf("failed to start mock redis: %v", err)
	}

	// 建立 Redis 客戶端連接到模擬的 Redis 伺服器
	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return rdb
}
