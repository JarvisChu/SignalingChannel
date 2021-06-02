package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/signalingchannel/ws"
)

func main() {
	r := gin.Default()
	r.GET("/ws/p2p", ws.HandleP2P)
	r.GET("/ws/room", ws.HandleRoom)
	r.Run(":8080")
}
