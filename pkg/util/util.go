package util

import (
	"io"
	"os"
	"time"
)

func CurrentDateTimeString() string {
	location := time.FixedZone("GMT+7", 7*3600)
	currentTime := time.Now().In(location)
	return currentTime.Format("2006-01-02 15:04:05")
}

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
