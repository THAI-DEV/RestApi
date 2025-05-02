package util

import (
	"log"
	"os"
	"time"
)

func ReadFile() (string, error) {
	data, err := os.ReadFile("./output/data.txt")
	if err != nil {
		log.Printf("failed reading data from file: %s \n", err)
		data = nil
	}

	log.Printf("File contents: %s", data)

	dataStr := string(data)
	return dataStr, err
}

func WriteFile(data string) error {
	// create output directory if not exists
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.Mkdir("./output", 0755)
	}

	err := os.WriteFile("./output/data.txt", []byte(data), 0644)
	if err != nil {
		log.Printf("failed writing data to file: %s \n", err)
		return err
	}

	return nil
}

func DiffDay(date1, date2 *time.Time) int {
	if date1 == nil || date2 == nil {
		return 0
	}
	diff := date2.Sub(*date1)
	return int(diff.Hours() / 24)
}
