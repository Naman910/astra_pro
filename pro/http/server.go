package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pro/service"
	"pro/utils"
)

const tempDir = "/tmp/astra"

func init() {
	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Message string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Unable to parse JSON request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	data := []byte(payload.Message)
	fileName := filepath.Join(tempDir, "data_"+utils.GenerateID()+".txt")
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		http.Error(w, "Unable to write data to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Data received"))

	// Launch a goroutine to process the data asynchronously
	go service.ProcessFile(fileName)
}
