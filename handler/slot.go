package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

// SlotHandler 處理slot遊戲的抽獎請求
func SlotHandler(db *gorm.DB, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// board := slot.GenerateBoard()
		// total, details := slot.CalculateWin(board, slot.Paylines, slot.Paytable, "Wild", "Scatter")

		// // 存 Redis
		// key := "slot:last_result"
		// data, _ := json.Marshal(details)
		// rdb.Set(config.Ctx, key, string(data), 0)

		// // 存 MySQL
		// boardJSON, _ := json.Marshal(board)
		// detailJSON, _ := json.Marshal(details)
		// db.Create(&model.SlotResult{
		// 	UserID:  "user123",
		// 	Board:   string(boardJSON),
		// 	Total:   total,
		// 	Details: string(detailJSON),
		// })

		// c.JSON(200, gin.H{
		// 	"board":   board,
		// 	"total":   total,
		// 	"details": details,
		// })
	}
}
