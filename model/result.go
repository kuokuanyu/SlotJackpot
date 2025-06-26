package model

type SlotResult struct {
	ID     uint `gorm:"primaryKey"`
	UserID string
	Board  string
	Total  int
	// Details string
}

// 指定對應的資料表名稱
func (SlotResult) TableName() string {
	return "slot_results"
}
