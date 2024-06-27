package db

import (
	"database/sql"
	"log"
	"sync"

	"pro/utils"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var (
	db   *sql.DB
	dbMu sync.Mutex
)

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	// Verify the connection is valid
	return db.Ping()
}

func StoreInDatabase(data []byte) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	id := utils.GenerateID()
	query := "INSERT INTO your_table_name (id, data) VALUES (?, ?)"
	_, err := db.Exec(query, id, data)
	if err != nil {
		log.Printf("Failed to store data: %v", err)
		return err
	}

	log.Printf("Data stored in database with ID: %s", id)
	return nil
}
