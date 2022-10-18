package web

import (
	"ChatSocket/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

func StartRoutingServerAPI() {
	router := gin.Default()
	router.GET("/", api.RootEndpoint)
	err := router.Run("localhost:8000")

	if err != nil {
		fmt.Println("Error up server with api ", err)
	} else {
		fmt.Println("Server with API upped")
	}
}

func StartWorkServerSockets() {
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
		go ProcessHandlingClient(conn)
	}
	fmt.Printf("Подключение прошло успешно по адресу %v", listen)
}

func StartServers() {
	StartRoutingServerAPI()
	StartWorkServerSockets()
}
