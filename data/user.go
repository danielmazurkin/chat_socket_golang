package data

import (
	"database/sql"
	"fmt"
)

type UserStruct struct {
	Username string `json:username`
	Password string `json:password`
}

func (*UserStruct) CreateTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS users(
			   ID INT PRIMARY KEY     NOT NULL,
			   NAME           TEXT    NOT NULL,
			   PASSWORD        CHAR(50)
			);`,
	)

	if err != nil {
		fmt.Println("Error with creating table", err)
	} else {
		fmt.Println("Table users was created")
	}
}

func (*UserStruct) CreateTableEntity(Username string, Password string) {

}
