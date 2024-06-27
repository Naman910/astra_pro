package service

import (
	"log"
	"os"
	"time"

	"pro/db"
)

func ProcessFile(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("Unable to read file %s: %v", fileName, err)
		return
	}

	// Simulate processing time
	time.Sleep(2 * time.Second)

	// Store data in database
	err = db.StoreInDatabase(data)
	if err != nil {
		log.Printf("Failed to store data from file %s: %v", fileName, err)
		return
	}

	// Remove the file after processing
	err = os.Remove(fileName)
	if err != nil {
		log.Printf("Failed to remove file %s: %v", fileName, err)
	}
}
