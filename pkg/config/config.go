package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var StartDateTime string
var Key string
var Mode string

func init() {
	readEnv()
}

func readEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Environment file (.env) not found, so using environment variables as a replacement")
	}

	Key = os.Getenv("KEY")
	Mode = os.Getenv("MODE")
	StartDateTime = time.Now().In(time.FixedZone("GMT+7", 7*3600)).Format("2006-01-02 15:04:05")
}
