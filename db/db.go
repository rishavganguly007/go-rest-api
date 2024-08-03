package db

import (
	"fmt"
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // _ tells its a sub-dependency
)

var DB *sql.DB
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("error creating DB")
	}

	DB.SetMaxOpenConns(10) // to create 10 pools
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTable  := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	)
	`
	DB.Exec(createUsersTable)
	_, err := DB.Exec(createUsersTable)

	createEventsTable  := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERECES users(id)
	)
	`

	DB.Exec(createEventsTable)

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic(fmt.Sprintf("could not create events table: %v", err))
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERECES users(id)
	FOREIGN KEY(event_id) REFERECES events(id)
	)
	
	`
	DB.Exec(createRegistrationTable)

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic(fmt.Sprintf("could not create registration table: %v", err))
	}
}