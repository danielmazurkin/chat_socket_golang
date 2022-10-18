package web

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ProcessHandlingClient(conn net.Conn) {

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("Ошибка выхода клиента = % v \n", err)
			return
		}

		fmt.Print(string(buf[:n]))
	}
}

func StartWorkClient() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("Ошибка подключения клиента: ", err)
		return
	}
	fmt.Println("Соединение успешно: ", conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		// Считываем строку пользовательского ввода с терминала и готовимся к отправке на сервер
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения строки", err)
		}

		line = strings.Trim(line, "\r\n")

		if line == "exit" {
			fmt.Println("Клиент выходит ..")
			break
		}

		// Отправляем строку на сервер
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write=", err)
		}
	}
}
