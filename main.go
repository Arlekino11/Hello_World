package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/greet", greetHendler)
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

func greetHendler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Anonymous"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")

	response := HealthResponse{
		Status:    "healty",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}
