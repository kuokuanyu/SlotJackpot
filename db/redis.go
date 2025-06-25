package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// time.Sleep(30 * time.Second)
	log.Println("🔗 正在連線到 Redis...")

	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Redis 連線失敗: %v", err)
	}

	log.Println("✅ Redis 已連線")
}
