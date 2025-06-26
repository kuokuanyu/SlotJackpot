package slot

// Paylines 定義玩家獲勝「連線」的規則線條。
// 定義哪些符號排列位置構成有效連線（可得分）
// 每一條線由 5 個座標組成 [column, row]
// 高度彈性	可以精確定義任何連線路徑（不侷限直線），適合商業 slot 的複雜設計需求。
// 符合真實機台邏輯	真實 slot 遊戲有很多 zigzag、diagonal、V 形、反 V 形等 Payline，這種 [列, 行] 設計能完全表達。
// 以下是一個較符合商業邏輯的 Paylines 陣列設計，包含：
// 三條水平線（上中下）
// 兩條對角線（左上到右下、左下到右上）
// 多條 Zigzag（Z字形／W字形）斜線
var Paylines = [][][2]int{
	// 1. Top Row
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, // 上排直線

	// 2. Middle Row
	{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}}, // 中排直線

	// 3. Bottom Row
	{{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}}, // 下排直線

	// 4. Diagonal Left Top → Right Bottom
	{{0, 0}, {1, 1}, {2, 2}, {3, 1}, {4, 0}}, // V字型對角

	// 5. Diagonal Left Bottom → Right Top
	{{0, 2}, {1, 1}, {2, 0}, {3, 1}, {4, 2}}, // 倒V字型對角

	// 6. Zigzag Top → Mid → Top → Mid → Top
	{{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 0}}, // Z字型

	// 7. Zigzag Bottom → Mid → Bottom → Mid → Bottom
	{{0, 2}, {1, 1}, {2, 2}, {3, 1}, {4, 2}}, // 倒Z字型

	// 8. W Shape
	{{0, 0}, {1, 2}, {2, 1}, {3, 2}, {4, 0}}, // W字型

	// 9. M Shape
	{{0, 2}, {1, 0}, {2, 1}, {3, 0}, {4, 2}}, // M字型

	// 10. Outward Arrow
	{{0, 1}, {1, 0}, {2, 1}, {3, 2}, {4, 1}}, // 箭頭展開形

	// 11. Inward Arrow
	{{0, 1}, {1, 2}, {2, 1}, {3, 0}, {4, 1}}, // 箭頭收束形
}

// CalculateWin 根據盤面計算中獎金額
// board: spin 出來的 5x3 盤面
// paylines: 所有連線的座標組合
// paytable: 各符號對應不同連線長度的分數
// 返回總贏分和每條連線得分詳細
func CalculateWin(board [][]string, paylines [][][2]int, paytable map[string]map[int]int, wildSymbol, scatterSymbol string) (int, map[int]int) {
	totalWin := 0
	lineWins := make(map[int]int) // 紀錄每條 payline 的贏分，key: payline index

	// 計算每條 payline 是否中獎
	for idx, line := range paylines {
		firstSymbol := ""   // 第一格的符號（用來比對連線）
		matchCount := 0     // 連續符號數量
		isWildMatch := true // 初始假設都是wild(可替代所有)

		for i, pos := range line {
			col, row := pos[0], pos[1] // 取得每個位置的列跟行
			symbol := board[col][row]  // 取得該行列的符號

			// 忽略scatter計算連線，scatter不需連線
			// 最後直接計算scatter數量
			if symbol == scatterSymbol {
				break
			}

			if i == 0 {
				// 第一格，直接設定 firstSymbol
				firstSymbol = symbol

				// 如果第一格是wild，先不定義firstSymbol，等待下一個非wild
				if firstSymbol == wildSymbol {
					isWildMatch = true
				} else {
					// 如果第一格不是wild，則確定firstSymbol
					isWildMatch = false
				}

				// 如果是第一格，則開始計算連線
				matchCount = 1
			} else {
				// 判斷是否符號相同 或 是wild可以替代
				if symbol == firstSymbol || // 符號相同
					symbol == wildSymbol || // wild，可以替代其他符號
					(firstSymbol == wildSymbol && symbol != scatterSymbol) { // 如果第一格是wild，且現在符號不是scatter，則視為連線
					matchCount++ // 累加連線數量

					// 如果第一格是wild，且現在符號不是wild，更新firstSymbol，確定連線符號
					if isWildMatch && symbol != wildSymbol {
						firstSymbol = symbol // 更新 firstSymbol 為當前符號
						isWildMatch = false  // 標記為非wild連線
					} else if isWildMatch && symbol == wildSymbol {
						// 如果第一格是wild，且現在也是wild，則繼續視為連線
						// 這樣可以允許連線中有多個wild
					} else if !isWildMatch && symbol == wildSymbol {
						// 如果第一格不是wild，但現在是wild，則視為連線
						isWildMatch = true
					} else if !isWildMatch && symbol != firstSymbol {
						// 如果第一格不是wild，但現在符號不同於firstSymbol，則中斷連線計算
						break
					}
				} else {
					// 如果符號不相同，則中斷連線計算
					break
				}
			}
		}

		// 最少連線3個符號才有獎金
		if matchCount >= 3 {
			// log.Println("matchCount: ", matchCount)

			// 從 paytable 取分數
			scoreMap, ok := paytable[firstSymbol] // firstSymbol 是連線的符號

			if ok {
				win := scoreMap[matchCount] // matchCount 是連線的數量
				totalWin += win             // 累加總分
				lineWins[idx] = win         // 記錄這條 payline 的獲得分數
			}
		}
	}

	// 計算 scatter 數量（scatter 不需要連線）
	scatterCount := 0
	for col := 0; col < len(board); col++ { // 0-4
		for row := 0; row < len(board[col]); row++ { // 0-2
			if board[col][row] == scatterSymbol {
				scatterCount++ // 累加 scatter 數量
			}
		}
	}
	if scatterCount >= 3 { // 如果有3個或以上的 scatter
		if scoreMap, ok := paytable[scatterSymbol]; ok {
			// 如果 paytable 有定義 scatter 獎金
			// 目前沒有定義 scatter 獎金
			if win, exists := scoreMap[scatterCount]; exists {
				totalWin += win
				lineWins[-1] = win // -1 代表 scatter 獎金，不在 payline 中
			}
		} else {
			// 如果 paytable 沒有定義 scatter 獎金，則不計算
			// 這裡可以根據實際需求決定是否要給 scatter 獎金
			// 例如：如果 scatter 是用來觸發免費遊戲或其他獎勵，而不是直接得分
		}
	}

	return totalWin, lineWins
}
