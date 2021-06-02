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

func (u *User)Send(messageType int, data []byte){
	u.mtx.Lock()
	defer u.mtx.Unlock()

	if u.Conn != nil {
		
	}
}