package data

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
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
		fmt.Println("Error with creating table", err)
	} else {
		fmt.Println("Table users was created")
	}
}

func (*UserStruct) CreateTableEntity(db *sql.DB, username string, password string) {
	sqlStatement := `INSERT INTO users('name', 'password') VALUES ($1, $2);`
	passwordForDatabase := GetMD5Hash(password)
	_, err := db.Exec(
		sqlStatement, username, passwordForDatabase,
	)

	if err != nil {
		fmt.Println("Error creating user in chat")
	} else {
		fmt.Println("User was created succesfull")
	}
}

func (*UserStruct) HasEntityTable(db *sql.DB, username string) {
	sqlStatement := `SELECT 1 FROM users WHERE name = $1;`

	result, err := db.Exec(sqlStatement, username)

	if err != nil {
		fmt.Println("Error checking username: ", err)
	} else {
		fmt.Println("Result executing query: ", result)
	}

}
