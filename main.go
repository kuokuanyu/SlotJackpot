package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kuokuanyu/SlotJackpot/db"
)

// build容器: docker build -t a167829435/SlotJackpot-app:latest .
// 啟動容器: docker run -d --name SlotJackpot-container  --env-file .env -p 8080:8080 a167829435/SlotJackpot-app:latest
// 停止容器: docker stop SlotJackpot-container
// 移除容器: docker rm SlotJackpot-container
func init() {
	log.Println("🔧 初始化應用程式...")

	// 本地開發時才載入 .env
	_ = godotenv.Load(".env")
}

func main() {

	// 初始化 MySQL 與 Redis
	db.InitMySQL()
	db.InitRedis()

	// 設置 API 路由
	r := gin.Default()
	// r.GET("/ping", handler.Ping)
	// 其他 API ...

	port := os.Getenv("PORT")
	r.Run(":" + port) // 監聽 PORT
}
