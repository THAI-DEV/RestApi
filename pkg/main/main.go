package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	exec()
}

func exec() {
	// Create a new gin instance
	router := gin.Default()

	// Define a route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
			"access":  time.Now().Format("2006-01-02 15:04:05"),
			"key":     readEnv(),
		})
	})

	// Run the server
	router.Run(":8080") // Default listens and serves on 0.0.0.0:8080
}

func readEnv() string {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	key := os.Getenv("KEY")
	return key

}
