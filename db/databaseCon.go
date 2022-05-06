package db

import (
	"database/sql"
	"fmt"
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
	// result, err := databasePostgres.Exec(
	// 	"INSERT INTO users(username, first_name, last_name, email, password, birthday) VALUES ($1, $2, $3, $4, $5, $6);",
	// 	"Temir009kz",
	// 	"Temirbolat 2",
	// 	"Maratuly 2",
	// 	"t_maratuly@gmail.com",
	// 	"12345",
	// 	"2001-01-31",
	// )
	result, err := databasePostgres.Exec("call add_friend();")
	if err != nil {
		log.Fatalf("During INSERT command appeared error: %v", err)
	}
	fmt.Println(result.RowsAffected())
	fmt.Println(result)
	return databasePostgres
}
