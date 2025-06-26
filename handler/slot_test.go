package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kuokuanyu/SlotJackpot/db"
)

// SlotHandler 測試
// 這個測試會模擬一個 GET 請求到 /slot 路由，並檢查回應是否符合預期
// 注意：這個測試需要依賴於 db.InitMockMySQL() 和 db.InitMockRedis()
// 這兩個函數應該返回一個模擬的 MySQL 和 Redis 客戶端
// 這樣可以避免實際連接到資料庫
func TestSlotHandler(t *testing.T) {
	// 初始化假的 MySQL & Redis
	mysql := db.InitMockMySQL()
	redis := db.InitMockRedis()

	router := gin.Default()
	router.GET("/slot", SlotHandler(mysql, redis))

	req, _ := http.NewRequest("GET", "/slot", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Expected status 200 but got %d", w.Code)
	}

	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)

	if res["total"] == nil {
		t.Errorf("Expected 'total' in response")
	}
}
