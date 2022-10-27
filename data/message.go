package data

import (
	"database/sql"
	"fmt"
)

type MessageStruct struct {
	message string `json:message`
	id_user string `json:id_user`
}

func (*MessageStruct) CreateTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE message
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			message TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id)  REFERENCES companies (id)
		);`,
	)

	if err != nil {
		fmt.Println("Error creating message table: ", err)
	} else {
		fmt.Println("Table messages was created")
	}
}

func (*MessageStruct) CreateTableEntity(db *sql.DB, message string, user_id int) {
	_, err := db.Exec(
		`INSERT INTO message VALUES(?, ?)`,
		message, user_id,
	)

	if err != nil {
		fmt.Println("Error was save message to database ", err)
		return
	}
}
