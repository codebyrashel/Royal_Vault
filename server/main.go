package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	defaultPort = "8080"
)

func main() {
	http.HandleFunc("/health", healthHandler)

	addr := ":" + defaultPort
	fmt.Printf("Server is running on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}