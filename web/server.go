package web

import (
	"ChatSocket/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

func ProcessHandlingClient(conn net.Conn) {

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("Error exit client = % v \n", err)
			return
		}

		fmt.Print(string(buf[:n]))
	}
}

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
	fmt.Println("Server begin listening ...")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("Error listening =", err)
		return
	}

	defer listen.Close()

	for {
		fmt.Println("Waiting connected client")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Happen next error: ", err)
		} else {
			fmt.Printf("Accept () succesfful connection = %v by ip addr client = %v \n", conn, conn.RemoteAddr().String())
		}
		go ProcessHandlingClient(conn)
	}
	fmt.Printf("Connection was successful %v", listen)
}
