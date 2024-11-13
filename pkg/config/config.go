package config

import (
	"dechdev/pkg/util"
	"log"
	"os"

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
	StartDateTime = util.CurrentDateTimeString()
}
