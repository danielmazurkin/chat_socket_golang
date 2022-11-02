package data

import (
	"ChatSocket/logger"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

type UserStruct struct {
	Username string `json:username`
	Password string `json:password`
}

func (*UserStruct) CreateTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS users(
			   ID INTEGER PRIMARY KEY AUTOINCREMENT,
			   NAME           TEXT    NOT NULL,
			   PASSWORD        CHAR(50)
			);`,
	)

	if err != nil {
		logger.Log.Println("Error with creating table", err)
	} else {
		logger.Log.Println("Table users was created")
	}
}

func (*UserStruct) CreateTableEntity(db *sql.DB, username string, password string) {
	sqlStatement := `INSERT INTO users('name', 'password') VALUES ($1, $2);`
	passwordForDatabase := GetMD5Hash(password)
	_, err := db.Exec(
		sqlStatement, username, passwordForDatabase,
	)

	if err != nil {
		logger.Log.Println("Error creating user in chat", err)
	} else {
		logger.Log.Println("User was created succesfull")
	}
}

func (*UserStruct) HasEntityTable(db *sql.DB, username string) bool {
	var has_entity bool
	sqlStatement := `SELECT 1 FROM users WHERE name = $1;`

	if err := db.QueryRow(sqlStatement, username).Scan(&has_entity); err != nil {
		logger.Log.Println("Error checking username: ", err)
	} else {
		logger.Log.Println("Result executing query: ", has_entity)
	}

	return has_entity
}
