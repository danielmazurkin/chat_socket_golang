package main

import (
	"ChatSocket/data"
	"ChatSocket/web"
)

func main() {
	data.OpenDatabase()
	data.InitRedisStorage()
	web.StartRoutingServerAPI()
}
