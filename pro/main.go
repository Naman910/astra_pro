package main

import (
	"log"
	"net/http"

	"pro/db"
	customhttp "pro/http"
)

func main() {
	// Initialize the database connection
	err := db.InitDB("user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	http.HandleFunc("/submit", customhttp.SubmitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
