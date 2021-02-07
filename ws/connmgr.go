package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jarvischu/signalchannel/account"
	"sync"
)

var connMgr *ConnMgr

// ConnMgr 管理所有的连接
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

func (c *ConnMgr) AddConn(uid string, conn *websocket.Conn) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// 互踢逻辑
	connStored, ok := c.connMap[uid]
	if ok {
		connStored.Close()
	}

	// 加入到mgr
	c.connMap[uid] = conn
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

func (c *ConnMgr) HandleConn(uid string, conn *websocket.Conn) {
	fmt.Printf("[HandleConn] uid:%v, conn:%v \n", uid, conn.RemoteAddr().String())

	c.AddConn(uid, conn)
	account.AddAccount(&account.Account{
		ID:     uid,
		Status: account.Online,
	})

	// 读取数据
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read message failed, uid:%v, err:%v \n", uid, err)
			c.RemoveConn(uid)
			account.UpdateAccountStatus(uid, account.Offline)
			return
		}

		fmt.Printf("recieve message from %v, msgType:%v, msg:%v \n", uid, msgType, string(msg))
	}
}

func (c *ConnMgr) Send(from string, to string, data string) error {
	return nil
}
