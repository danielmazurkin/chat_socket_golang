package data

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(db *sql.DB) {
	user := new(UserStruct)
	user.CreateTableUser(db)
}

func OpenDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "database.db")

	if err == nil {
		fmt.Println("Connecting with database")
		InitDatabase(db)
	} else {
		fmt.Println("Error with connecting ", err)
	}

	return db, err
}
