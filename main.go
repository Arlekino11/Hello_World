package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type InfoResponse struct {
	Author    string   `json:"author"`
	Version   string   `json:"version"`
	Endpoints []string `json:"endpoints"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func main() {
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Anonymous"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := InfoResponse{
		Author:    "Arlekino",
		Version:   "1.0.0",
		Endpoints: []string{"/", "/health", "/greet", "/info"},
	}
	json.NewEncoder(w).Encode(response)
}
