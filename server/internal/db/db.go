package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase() (*Database, error) {
	dsn := getDSN()
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{Conn: conn}, nil
}

func getDSN() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5433") // use 5433 because Docker maps host 5433 -> container 5432
	user := getEnv("DB_USER", "royalvault")
	password := getEnv("DB_PASSWORD", "royalvault")
	dbname := getEnv("DB_NAME", "royal_vault")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// DSN format for lib/pq
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}