package main

import (
	"social-analyzer-backend/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	database.ConnectDB()

	// Run migrations
	database.Migrate()

	// Setup Gin
	r := gin.Default()

	// Disable trusted proxies warning (development only)
	r.SetTrustedProxies(nil)

	// CORS (contoh simple)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Routes (contoh sementara)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running"})
	})

	// Start server
	port := ":8080"
	fmt.Println("✅ Server running on http://localhost" + port)
	if err := r.Run(port); err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
