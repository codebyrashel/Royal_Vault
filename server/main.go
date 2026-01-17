package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codebyrashel/Royal_Vault/server/internal/db"
	"github.com/codebyrashel/Royal_Vault/server/internal/routes"
)

func main() {
	port := getPort()

	// Initialize database connection
	database, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Conn.Close()
	fmt.Println("Connected to database successfully")

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