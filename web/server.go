package web

import (
	"ChatSocket/api"
	"ChatSocket/logger"
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
			logger.Log.Println("Error exit client = % v \n", err)
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
		logger.Log.Println("Error up server with api ", err)
	} else {
		logger.Log.Println("Server with API upped")
	}
}

func StartWorkServerSockets() {
	logger.Log.Println("Server begin listening ...")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		logger.Log.Println("Error listening =", err)
		return
	}

	defer listen.Close()

	for {
		logger.Log.Println("Waiting connected client")
		conn, err := listen.Accept()

		if err != nil {
			logger.Log.Println("Happen next error: ", err)
		} else {
			logger.Log.Println("Accept () succesfful connection = %v by ip addr client = %v \n", conn, conn.RemoteAddr().String())
		}
		go ProcessHandlingClient(conn)
	}
	logger.Log.Println("Connection was successful %v", listen)
}
