package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Key string
var Mode string

func init() {
	readEnv()
}

func readEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		// log.Fatalf("Error loading .env file")
		log.Println("Error loading .env file")
	}

	Key = os.Getenv("KEY")
	Mode = os.Getenv("MODE")
}
