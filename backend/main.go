package main

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()
	Openai()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))


	r.GET("/api/hello", func(c *gin.Context){
		c.JSON(200,gin.H{"message":"Hello from Go!"})
	})
	r.Run(":8080")
}

