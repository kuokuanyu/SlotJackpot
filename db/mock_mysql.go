package db

import (
	"log"

	"github.com/kuokuanyu/SlotJackpot/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitMockMySQL 回傳一個記憶體中的 SQLite DB，用來模擬 MySQL
func InitMockMySQL() *gorm.DB {
	// 使用 SQLite 的記憶體模式來模擬 MySQL
	// 注意：這只是模擬，實際上 SQLite 和 MySQL 在某些行為上可能有差異
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect mock MySQL: %v", err)
	}

	// 自動建表，模擬你的 SlotResult 表
	err = db.AutoMigrate(&model.SlotResult{})
	if err != nil {
		log.Fatalf("failed to migrate mock MySQL: %v", err)
	}

	return db
}
