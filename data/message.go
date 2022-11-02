package data

import (
	"ChatSocket/logger"
	"database/sql"
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
			is_send BIT NOT NULL DEFAULT 0,
			FOREIGN KEY (user_id)  REFERENCES companies (id)
		);`,
	)

	if err != nil {
		logger.Log.Println("Error creating message table: ", err)
	} else {
		logger.Log.Println("Table messages was created")
	}
}

func (*MessageStruct) CreateTableEntity(db *sql.DB, message string, user_id int) {
	// 0 is mark as message not read
	_, err := db.Exec(
		`INSERT INTO message VALUES(?, ?, ?)`,
		message, user_id, 0,
	)

	if err != nil {
		logger.Log.Println("Error was save message to database ", err)
		return
	}
}

func (*MessageStruct) HasEntityTable(db *sql.DB, username string) bool {
	var has_entity bool
	sqlStatement := `SELECT 1 FROM message WHERE username = $1;`

	if err := db.QueryRow(sqlStatement, username).Scan(&has_entity); err != nil {
		logger.Log.Println("Error checking username: ", err)
	} else {
		logger.Log.Println("Result executing query: ", has_entity)
	}

	return has_entity
}

func (*MessageStruct) UpdateMessageMarkAsSend(db *sql.DB, username string) bool {
	return false
}
