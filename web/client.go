package web

import (
	"ChatSocket/data"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartWorkClient() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	db, err := data.OpenDatabase()
	if err != nil {
		fmt.Println("Error connecting client: ", err)
		return
	}
	fmt.Println("Connecting successful: ", conn)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("1) Register in chat")
	fmt.Println("2) Enter in chat")

	line, err := reader.ReadString('\n')
	line = strings.Trim(line, "\r\n")

	if line == string('1') {
		ClientRegistration(db)
	} else if line == string('2') {
		username := EnterInChat(db)

		if len(username) == 0 {
			fmt.Println("User not created..")
			return
		}

	}

	for {
		fmt.Println("Enter message to chat..")
		// Считываем строку пользовательского ввода с терминала и готовимся к отправке на сервер
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error read string", err)
		}

		line = strings.Trim(line, "\r\n")

		if line == "exit" {
			fmt.Println("Client exit ..")
			break
		}

		// Отправляем строку на сервер
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write=", err)
		}
	}

	defer db.Close()
}
