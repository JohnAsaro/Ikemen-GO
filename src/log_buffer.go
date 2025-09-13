package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func logScreenBuffer() {

	// Create/open database
	path := fmt.Sprintf("external/mods/bridges/bridge_%s.db", sys.instanceID)
	db, err := sql.Open("sqlite3", path)
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
	updateSQL := `UPDATE buffer SET width = ?, height = ?, buffer_data = ?, done = 0 
				WHERE id = (SELECT id FROM buffer WHERE done = -1 ORDER BY id LIMIT 1)`
	_, err = db.Exec(updateSQL, width, height, pixdata)
	if err != nil {
		log.Println("Update failed:", err)
		return
	}
}
