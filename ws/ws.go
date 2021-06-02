package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func HandleP2P(c *gin.Context) {

	// check params
	name := c.Query("name")
	if len(name) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	// upgrade to websocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	GetP2PMgr().UserLogin(&User{
		Name: name,
		Conn: conn,
	})

	return
}

func HandleRoom(c *gin.Context) {

	// check params
	name := c.Query("name")
	if len(name) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	roomID := c.Query("roomid")
	if len(roomID) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	// upgrade to websocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	GetRoomMgr().UserEnterRoom(&User{
		Name: name,
		Conn: conn,
	}, roomID)

	return
}