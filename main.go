package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/signalchannel/api"
	"github.com/jarvischu/signalchannel/ws"
)

func main() {
	r := gin.Default()
	r.GET("/ping", api.Ping)
	r.GET("/accounts", api.GetAccountList)
	r.GET("/ws", ws.Handle)
	r.Run(":8080")
}
