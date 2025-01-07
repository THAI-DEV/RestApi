package config

import (
	"log"
	"os"

	"github.com/THAI-DEV/dechutil"
	"github.com/joho/godotenv"
)

var StartTime string
var BuildTime string
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
	StartTime = dechutil.TimeToStringDateTimeFull(dechutil.CurrentBkkTime())
}
