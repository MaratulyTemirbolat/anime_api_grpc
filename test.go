package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type ResultMy struct {
	id        uint32
	firstName string
	lastName  string
	email     string
	username  string
	phones    []string
	isBlocked bool
}

// Не забыть в конце закрыть Базу данных!
func main() {
	var connectionDB string = "user = postgres password = 1538 dbname = trial_grpc sslmode=disable"
	databasePostgres, err := sql.Open("postgres", connectionDB)
	if err != nil {
		log.Fatalf("Cannot connect to db by general settings, error: %v", err)
	}
	defer databasePostgres.Close()
	// result, err := databasePostgres.Exec(
	// 	"INSERT INTO users(username, first_name, last_name, email, password, birthday) VALUES ($1, $2, $3, $4, $5, $6);",
	// 	"Temir009kz",
	// 	"Temirbolat 2",
	// 	"Maratuly 2",
	// 	"t_maratuly@gmail.com",
	// 	"12345",
	// 	"2001-01-31",
	// )
	var isBlockedU bool
	isBlockRow, err := databasePostgres.Query("SELECT is_blocked FROM friends WHERE user_a_id = 2 AND user_b_id = 1 LIMIT 1;")
	if err != nil {
		log.Fatalf("Appeared BOOL error: %v", err)
	}
	for isBlockRow.Next() {
		err := isBlockRow.Scan(&isBlockedU)
		if err != nil {
			log.Fatal("ERROR BOOL: ", err)
			continue
		}
	}
	resultRow, err := databasePostgres.Query(
		"SELECT users.id, users.first_name, users.last_name, users.email, users.username, phones.phone FROM users LEFT JOIN phones ON users.id = phones.owner_id WHERE users.id = 1;")
	if err != nil {
		log.Fatalf("Appeared error: %v", err)
	}
	defer resultRow.Close()
	user := []ResultMy{}

	for resultRow.Next() {
		us := ResultMy{}
		var phonee string
		err := resultRow.Scan(&us.id, &us.firstName, &us.lastName, &us.email, &us.username, &phonee)
		us.phones = append(us.phones, phonee)
		if err != nil {
			log.Fatal("ERROR: ", err)
			continue
		}
		user = append(user, us)
	}
	resultedUser := ResultMy{}
	if len(user) != 0 {
		resultedUser.id = user[0].id
		resultedUser.firstName = user[0].firstName
		resultedUser.lastName = user[0].lastName
		resultedUser.email = user[0].email
		resultedUser.username = user[0].username
		resultedUser.isBlocked = isBlockedU
		for _, v := range user {
			resultedUser.phones = append(resultedUser.phones, v.phones[0])
		}

		log.Fatal("Your user: ", resultedUser)
	}
	log.Fatal("USER WITH THIS ID IS NOT FOUND")
}
