package web

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

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

func StartWorkServer() {
	fmt.Println("Сервер начал прослушивание ...")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("Ошибка прослушивания =", err)
		return
	}

	defer listen.Close()

	for {
		fmt.Println("Ожидание подключения клиента")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Возникла следующая ошибка: ", err)
		} else {
			fmt.Printf("Accept () успешное соединение = %v по ip клиента = %v \n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
	fmt.Printf("Подключение прошло успешно по адресу %v", listen)
}
