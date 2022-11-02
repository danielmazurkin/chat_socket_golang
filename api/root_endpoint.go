package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RootEndpoint(c *gin.Context) {
	msg := "Hello in future is will be chat in golang"
	_, err := json.Marshal(msg)

	if err != nil {
		log.Println("Error with decoding")
		return
	}

	c.JSON(http.StatusOK, msg)
}
