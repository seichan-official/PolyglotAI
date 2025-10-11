package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/seichan-official/polyglotai-backend/internal/service"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Go!"})
	})

	// 翻訳エンドポイント
	r.POST("/api/generate", func(c *gin.Context) {
		var req struct {
			Text string `json:"text"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		// Ollamaを使って翻訳
		prompt := fmt.Sprintf("以下の文章を英語に翻訳してください。翻訳結果のみを返してください：\n%s", req.Text)
		translatedText, err := service.CallOllama(c.Request.Context(), "llama3", prompt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Translation failed"})
			return
		}

		c.JSON(200, gin.H{"translatedText": translatedText})
	})

	// サーバーはゴルーチンで起動
	go func() {
		if err := r.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	// テスト用に Ollama 呼び出し
	time.Sleep(1 * time.Second) // サーバーが起動するまで少し待つ
	service.TestOllama()

	// メインをブロックさせる（任意）
	select {}
}
