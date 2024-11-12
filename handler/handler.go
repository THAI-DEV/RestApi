package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "-- pong --")
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
		"access":  time.Now().In(time.UTC).Format("2006-01-02 15:04:05"),
		"key":     readEnv(),
	})
}

func readEnv() string {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		// log.Fatalf("Error loading .env file")
		log.Println("Error loading .env file")
	}

	key := os.Getenv("KEY")
	return key

}
