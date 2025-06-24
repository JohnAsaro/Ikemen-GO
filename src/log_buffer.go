package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func logScreenBuffer() {

	// Create/open database
	db, err := sql.Open("sqlite3", "external/mods/bridge.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	createTable := `
	CREATE TABLE IF NOT EXISTS buffer (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		width INTEGER NOT NULL,
		height INTEGER NOT NULL,
		buffer_data BLOB NOT NULL,
		done INTEGER NOT NULL DEFAULT 0
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	pixdata, width, height := captureScreenBuffer() // Get screen buffer

	// Insert into database
	insertSQL := `INSERT INTO buffer (width, height, buffer_data) VALUES (?, ?, ?)`
	_, err = db.Exec(insertSQL, width, height, pixdata)
	if err != nil {
		log.Fatal(err)
	}

}
