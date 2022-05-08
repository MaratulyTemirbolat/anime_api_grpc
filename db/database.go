package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Не забыть в конце закрыть Базу данных!
func ConnectGenDB() *sql.DB {
	var connectionDB string = "user = postgres password = 1538 dbname = trial_grpc sslmode=disable"
	databasePostgres, err := sql.Open("postgres", connectionDB)
	if err != nil {
		log.Fatalf("Cannot connect to db by general settings, error: %v", err)
	}

	return databasePostgres
}
