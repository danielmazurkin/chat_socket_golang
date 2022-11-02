package data

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

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
		log.Println("Error with creating table", err)
	} else {
		log.Println("Table users was created")
	}
}

func (*UserStruct) CreateTableEntity(db *sql.DB, username string, password string) {
	sqlStatement := `INSERT INTO users('name', 'password') VALUES ($1, $2);`
	passwordForDatabase := GetMD5Hash(password)
	_, err := db.Exec(
		sqlStatement, username, passwordForDatabase,
	)

	if err != nil {
		log.Println("Error creating user in chat", err)
	} else {
		log.Println("User was created succesfull")
	}
}

func (*UserStruct) HasEntityTable(db *sql.DB, username string) bool {
	var has_entity bool
	sqlStatement := `SELECT 1 FROM users WHERE name = $1;`

	if err := db.QueryRow(sqlStatement, username).Scan(&has_entity); err != nil {
		log.Println("Error checking username: ", err)
	} else {
		log.Println("Result executing query: ", has_entity)
	}

	return has_entity
}
