package slot

// 不同欄位長度（增加隨機性與控制中獎率）
// 符號權重不同（高價符號少見）
// 仍保留「連續三格抽出」的邏輯（真實 slot 大多使用）
// 通常為 30～100 個 symbol。
// 有些大型機種甚至會到 200 個 symbol，因為這樣可以精細控制中獎機率（Hit Rate）與賠率分佈。
// 每個 Reel 長度不同（48～63 個符號），模擬真實 slot 輪帶長度。
// A ~ 9	基本中低價值符號，常見組合
// Wild	百搭符號，可取代其他非特殊符號
// Scatter	散佈符號，觸發免費遊戲或額外獎勵（不限線）
// Bonus	獎勵符號，收集可觸發 Bonus Game
// Dragon	高分特殊符號，可能與 Tiger 組成特殊模式
// Tiger	同上，搭配 Dragon 觸發大獎或特殊玩法
// 更貼近實際場景
var Reels = [][]string{
	// Reel 1 - 52 個符號
	{"A", "10", "K", "Q", "J", "9", "Wild", "A", "Bonus", "K", "Q", "J", "Scatter",
		"A", "10", "Q", "9", "Wild", "Dragon", "K", "J", "9", "Q", "A", "Tiger",
		"10", "9", "J", "Bonus", "Q", "Wild", "K", "A", "Scatter", "10", "Q", "J",
		"K", "10", "9", "Wild", "Scatter", "A", "Dragon", "Q", "J", "10", "K", "Tiger"},

	// Reel 2 - 60 個符號
	{"Q", "A", "Wild", "9", "K", "J", "Scatter", "Bonus", "J", "Q", "Wild", "K", "A", "10",
		"9", "Q", "Scatter", "K", "J", "Q", "10", "9", "Wild", "Dragon", "A", "J", "Bonus",
		"Q", "K", "J", "10", "A", "Wild", "9", "K", "Q", "J", "A", "10", "9", "K",
		"Wild", "Q", "J", "Scatter", "A", "K", "Tiger", "9", "Wild", "Q", "A", "Bonus", "10", "K", "Dragon", "9", "J"},

	// Reel 3 - 48 個符號（中間短，含特殊符號較多）
	// Reel 3 較短，用來控制中間中獎機率，適合放特殊符號多一些。
	{"10", "9", "K", "Q", "A", "Scatter", "Wild", "J", "A", "Bonus", "Q", "Wild",
		"9", "J", "Dragon", "Q", "Scatter", "A", "Wild", "K", "Tiger", "Q", "10", "9",
		"Scatter", "K", "A", "J", "Q", "10", "Wild", "Bonus", "Dragon", "J", "Q", "A",
		"10", "9", "Wild", "K", "A", "Q", "J", "Tiger", "9", "Bonus"},

	// Reel 4 - 56 個符號
	{"A", "K", "Wild", "Scatter", "10", "9", "Q", "J", "K", "Bonus", "A", "Wild",
		"9", "J", "Q", "Scatter", "K", "Tiger", "10", "A", "Q", "Wild", "9", "K",
		"J", "Q", "10", "Scatter", "A", "Dragon", "Wild", "9", "10", "Q", "J", "A",
		"K", "10", "Wild", "Scatter", "J", "Bonus", "A", "10", "9", "K", "Q", "Wild",
		"10", "J", "Dragon", "A", "K", "Wild"},

	// Reel 5 - 63 個符號（最長，控制中獎率）
	// Reel 5 較長，讓右邊中線的機率更低或具控制性。
	{"Wild", "A", "Q", "K", "J", "Scatter", "10", "9", "Wild", "Q", "A", "J", "K", "Scatter",
		"9", "10", "K", "Wild", "J", "Bonus", "10", "9", "A", "K", "Wild", "Q", "J", "10",
		"Scatter", "A", "Dragon", "9", "Q", "J", "Wild", "10", "A", "Q", "K", "Scatter", "J", "9",
		"Wild", "A", "K", "Bonus", "Q", "J", "Scatter", "10", "9", "K", "A", "Wild", "Q", "Tiger",
		"10", "Scatter", "A", "K", "Q", "J", "Wild"},
}

// Reels 為每個轉輪的符號組成。
// 假設為 5 個轉輪，每個轉輪內可能出現的符號都一樣。
// var Reels = [][]string{
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
// }
