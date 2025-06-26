package slot

import "testing"

// TestCalculateWin 測試 CalculateWin 函數是否能正確計算中獎分數
func TestCalculateWin(t *testing.T) {
	board := [][]string{
		{"A", "A", "A"},
		{"K", "Q", "J"},
		{"Q", "Q", "Q"},
		{"K", "J", "10"},
		{"A", "A", "A"},
	}

	expected := 0 // 預期沒有中獎分數
	total, _ := CalculateWin(board, Paylines, Paytable, "Wild", "Scatter")
	if total != expected {
		t.Errorf("expected %d, got %d", expected, total)
	}
}
