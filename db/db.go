package db

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	db_instance, err := sql.Open("sqlite3", "gopherly.db")
	DB = db_instance

	if err != nil {
		panic("Cannot initialise database...")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Cannot create events table...")
	}
}
