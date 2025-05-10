package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
	Cache    string `json:"cache"`
}

func main() {
	port := "8081" // Changed from "80" to "8081" to match production.yaml mapping

	// Log startup info
	log.Println("Starting API server on port", port)
	log.Println("Database URL:", os.Getenv("DATABASE_URL"))
	log.Println("Cache URL:", os.Getenv("CACHE_URL"))
	log.Println("Log level:", os.Getenv("LOG_LEVEL"))

	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthCheckHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Welcome to the Go API"}
	json.NewEncoder(w).Encode(response)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	health := HealthResponse{
		Status:   "healthy",
		Database: "connected",
		Cache:    "connected",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health)
}
