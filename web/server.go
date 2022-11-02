package web

import (
	"ChatSocket/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func ProcessHandlingClient(conn net.Conn) {

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			log.Printf("Error exit client = % v \n", err)
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
		log.Println("Error up server with api ", err)
	} else {
		log.Println("Server with API upped")
	}
}

func StartWorkServerSockets() {
	log.Println("Server begin listening ...")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		log.Println("Error listening =", err)
		return
	}

	defer listen.Close()

	for {
		log.Println("Waiting connected client")
		conn, err := listen.Accept()

		if err != nil {
			log.Println("Happen next error: ", err)
		} else {
			log.Printf("Accept () succesfful connection = %v by ip addr client = %v \n", conn, conn.RemoteAddr().String())
		}
		go ProcessHandlingClient(conn)
	}
	log.Printf("Connection was successful %v", listen)
}
