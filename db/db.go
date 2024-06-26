package db

import (
	"database/sql"
	"github.com/hkm007/gopherly/utils/constants"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	db_instance, err := sql.Open("sqlite3", constants.DB_NAME)
	if err != nil {
		panic("Cannot initialise database...")
	}
	DB = db_instance

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Cannot create users table...")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Cannot create events table...")
	}
}
