# 使用官方 Golang 圖像作為建置階段
FROM golang:1.24-alpine AS builder

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 並下載依賴
COPY go.mod go.sum ./
RUN go mod download

# 複製所有原始碼
COPY . .

# 編譯 Go 應用（將編譯出來的 binary 命名為 app）
RUN go build -o SlotJackpot-app .

# 運行階段的最小映像
FROM alpine:latest
WORKDIR /root/

# 複製編譯好的 binary 和 .env
COPY --from=builder /app/SlotJackpot-app .

# 設定容器啟動時執行的指令
CMD ["./SlotJackpot-app"]
