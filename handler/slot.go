package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kuokuanyu/SlotJackpot/db"
	"github.com/kuokuanyu/SlotJackpot/model"
	"github.com/kuokuanyu/SlotJackpot/slot"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

var i = 1

// SlotHandler 處理slot遊戲的抽獎請求
func SlotHandler(mysql *gorm.DB, redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 抽獎，產生 5x3 的盤面
		board := slot.Spin()
		// 根據盤面計算中獎金額 board: spin 出來的 5x3 盤面 paylines: 所有連線的座標組合 paytable: 各符號對應不同連線長度的分數 返回總贏分和每條連線得分詳細
		total, details := slot.CalculateWin(board, slot.Paylines, slot.Paytable, "Wild", "Scatter")

		// 整理連線結果
		lineResults := []gin.H{}
		for line, win := range details {
			lineResults = append(lineResults, gin.H{
				"type":   "line",
				"line":   line + 1,
				"score":  win,
				"detail": fmt.Sprintf("連線 %d 獲勝 %d 分", line+1, win),
			})
		}

		// 存 Redis
		key := "slot:last_result" // 使用 i 作為唯一識別符
		data, _ := json.Marshal(details)                  // 將 details 轉換為 JSON 格式
		redis.Set(db.Ctx, key, string(data), 0)           // 寫入redis，0表示不過期

		// 存 MySQL
		boardJSON, _ := json.Marshal(board) // 將盤面轉換為 JSON 格式
		// detailJSON, _ := json.Marshal(details)
		mysql.Create(&model.SlotResult{
			UserID: "user123",
			Board:  string(boardJSON),
			Total:  total,
			// Details: string(detailJSON),
		})

		c.JSON(200, gin.H{
			// "board":   board,
			"total":     total,
			"win_lines": lineResults,
		})

		i++ // 增加 i 的值，確保每次請求都有唯一的 key
	}
}

// 將盤面轉換為可讀格式（row-major 3x5）
// displayBoard := make([][]string, 3) // 3 row
// for row := 0; row < 3; row++ {
// 	displayBoard[row] = make([]string, 5)
// 	for col := 0; col < 5; col++ {
// 		displayBoard[row][col] = board[col][row]
// 	}
// }
