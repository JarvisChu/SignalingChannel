package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/signalchannel/ws"
)

func HandleHTTPPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", HandleHTTPPing)
	r.GET("/ws", ws.Handle)
	r.Run(":8080")
}
