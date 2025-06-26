package slot

import (
	"math/rand"
)


// Spin 函式會隨機選取每個轉輪上的 3 個連續符號，組成一個 5x3 的遊戲盤面
func Spin() [][]string {
	board := make([][]string, 5) // 建立 5 個欄位（轉輪）

	for i := 0; i < 5; i++ {
		reel := Reels[i]              // 第 i 個轉輪的符號列表
		start := rand.Intn(len(reel)) // 隨機起始位置
		window := []string{           // 取出 3 個連續符號
			reel[(start+0)%len(reel)],
			reel[(start+1)%len(reel)],
			reel[(start+2)%len(reel)],
		}
		board[i] = window // 存入盤面
	}

	return board // 回傳盤面：5 x 3（欄 x 行）
}
