package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
	"sync"
)

var p2pMgr *P2PMgr

// P2PMgr manage all connections
type P2PMgr struct {
	users map[string]*User
	mtx     sync.Mutex
}

// GetP2PMgr
func GetP2PMgr() *P2PMgr {
	if p2pMgr == nil {
		p2pMgr = &P2PMgr{
			users: make(map[string]*User, 0),
		}
	}
	return p2pMgr
}

func (mgr *P2PMgr) addUser(user *User) {
	mgr.mtx.Lock()
	defer mgr.mtx.Unlock()
	mgr.users[user.Name] = user
}

func (mgr *P2PMgr) removeUser(userName string) {
	mgr.mtx.Lock()
	defer mgr.mtx.Unlock()
	delete(mgr.users, userName)
}

func (mgr *P2PMgr) UserLogin(user *User) {
	mgr.UserLogout(user.Name)
	mgr.addUser(user)
	mgr.handleConn(user)
}

func (mgr *P2PMgr) UserLogout(userName string){
	mgr.removeUser(userName)
}

func (mgr *P2PMgr) getConnByUserName(userName string) *websocket.Conn {
	user, ok := mgr.users[userName]
	if !ok {
		return nil
	}
	return user.Conn
}

func (mgr *P2PMgr) handleConn(user *User) {
	fmt.Printf("[handleConn] name:%v, conn:%v \n", user.Name, user.Conn.RemoteAddr().String())

	// Read data
	peer := ""
	for {
		msgType, msg, err := user.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read message failed, name:%v, err:%v \n", user.Name, err)

			// if disconnect by client
			mgr.UserLogout(user.Name)
			return
		}

		fmt.Printf("recieve message from %v, msgType:%v, msg:%v \n", user.Name, msgType, string(msg))

		// set-peer
		if strings.HasPrefix(string(msg), "set-peer:") {
			arr := strings.Split(string(msg), ":")
			if len(arr) == 2 {
				peer = arr[1]
			}
			continue
		}

		// forward message to peer
		if len(peer) > 0 {
			peerConn := mgr.getConnByUserName(peer)
			if peerConn != nil {
				if err := peerConn.WriteMessage(msgType,msg); err != nil {
					fmt.Printf("WriteMessage failed, %v\n", err)
				}else{
					fmt.Printf("WriteMessage success\n")
				}
			}
		}
	}
}