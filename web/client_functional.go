package web

import (
	"ChatSocket/data"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func ClientRegistration(db *sql.DB) {
	user := new(data.UserStruct)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your username")

	username, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading string: ", err)
	}

	username = strings.Trim(username, "\r\n")
	user.HasEntityTable(db, username)
	fmt.Println("Enter your password")

	password, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading string: ", err)
	}

	password = strings.Trim(password, "\r\n")
	user.CreateTableEntity(db, username, password)
}
