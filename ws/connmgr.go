package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jarvischu/signalchannel/account"
	"sync"
)

var connMgr *ConnMgr

// ConnMgr manage all connections
type ConnMgr struct {
	connMap map[string]*websocket.Conn
	mtx     sync.Mutex
}

// GetConnMgr
func GetConnMgr() *ConnMgr {
	if connMgr == nil {
		connMgr = &ConnMgr{
			connMap: make(map[string]*websocket.Conn, 0),
		}
	}
	return connMgr
}

func (c *ConnMgr) AddConn(id string, conn *websocket.Conn) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// disconnect previous connection
	connStored, ok := c.connMap[id]
	if ok {
		connStored.Close()
	}

	c.connMap[id] = conn
}

func (c *ConnMgr) RemoveConn(uid string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	conn, ok := c.connMap[uid]
	if ok {
		conn.Close()
	}

	delete(c.connMap, uid)
}

func (c *ConnMgr) GetConn(uid string) *websocket.Conn {
	return c.connMap[uid]
}

func (c *ConnMgr) HandleConn(uid string, conn *websocket.Conn) {
	fmt.Printf("[HandleConn] uid:%v, conn:%v \n", uid, conn.RemoteAddr().String())

	c.AddConn(uid, conn)
	account.AddAccount(&account.Account{
		ID:     uid,
		Status: account.Online,
	})

	// Send data
	go func() {
		//todo using channel to send data
	}()

	// Read data
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read message failed, uid:%v, err:%v \n", uid, err)

			// if disconnect by client, update connections and account
			connStored := c.GetConn(uid)
			if connStored != nil && conn != nil && connStored == conn {
				c.RemoveConn(uid)
				account.UpdateAccountStatus(uid, account.Offline)
			}
			return
		}

		fmt.Printf("recieve message from %v, msgType:%v, msg:%v \n", uid, msgType, string(msg))
	}
}

type DataFrame struct {
	Sender string `json:"sender"`
	Msg    string `json:"msg"`
}

func (c *ConnMgr) Send(from string, to string, msg string) error {
	conn := c.GetConn(to)
	if conn == nil {
		return fmt.Errorf("connection not found, to:%v", to)
	}

	frame := DataFrame{
		Sender: from,
		Msg:    msg,
	}

	b, _ := json.Marshal(&frame)
	return conn.WriteMessage(websocket.TextMessage, b)
}
