package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type User struct{
	Name string
	Conn *websocket.Conn
	mtx sync.Mutex
}

// gorilla websocket DO NOT support concurrent write
// so we need using mtx
func (u *User)Send(messageType int, data []byte){
	u.mtx.Lock()
	defer u.mtx.Unlock()

	if u.Conn != nil {
		u.Conn.WriteMessage(messageType, data)
	}
}