package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codebyrashel/Royal_Vault/server/internal/routes"
)

func main() {
	port := getPort()
	router := routes.SetupRouter()

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)

	if err := router.Run(addr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}