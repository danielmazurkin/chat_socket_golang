package data

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "database.db")

	if err == nil {
		fmt.Println("Connecting with database")
	} else {
		fmt.Println("Error with connecting ", err)
	}

	return db, err
}
