package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/signalingchannel/api"
	"github.com/jarvischu/signalingchannel/ws"
)

func main() {
	r := gin.Default()
	r.GET("/ping", api.Ping)
	r.GET("/accounts", api.GetAccountList)
	r.GET("/ws", ws.Handle)
	r.POST("/msg", api.SendMsg)
	r.Run(":8080")
}
