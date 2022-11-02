package web

import (
	"ChatSocket/data"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func ClientRegistration(db *sql.DB) {
	user := new(data.UserStruct)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your username")

	username, err := reader.ReadString('\n')

	if err != nil {
		log.Println("Error reading string: ", err)
	}

	username = strings.Trim(username, "\r\n")
	fmt.Println("Enter your password")

	password, err := reader.ReadString('\n')

	if err != nil {
		log.Println("Error reading string: ", err)
	}

	password = strings.Trim(password, "\r\n")
	user.CreateTableEntity(db, username, password)
}

func EnterInChat(db *sql.DB) string {
	user := new(data.UserStruct)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your username")

	username, err := reader.ReadString('\n')

	if err != nil {
		log.Println("Error reading string: ", err)
	}

	username = strings.Trim(username, "\r\n")
	has_user := user.HasEntityTable(db, username)

	var result string

	if has_user {
		result = username
	}
	log.Println("Result checking username ", result)
	return result
}
