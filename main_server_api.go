package main

import (
	"ChatSocket/data"
	"ChatSocket/web"
	"log"
)

func main() {
	_, err := data.OpenDatabase()

	if err != nil {
		log.Println("Error connecting client: ", err)
		return
	}

	rdb := data.InitRedisStorage()

	if rdb == nil {
		log.Println("Error connecting redis")
	}

	web.StartRoutingServerAPI()
}
