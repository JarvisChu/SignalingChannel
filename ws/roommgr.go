package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

var roomMgr *RoomMgr

type Room struct{
	ID string
	Users map[string]*User
}

type RoomMgr struct {
	mtx sync.Mutex
	rooms map[string]*Room
}

func GetRoomMgr() *RoomMgr {
	if roomMgr == nil {
		roomMgr = &RoomMgr{
			rooms: make(map[string]*Room, 0),
		}
	}
	return roomMgr
}

func (mgr *RoomMgr) broadCastMsg(room *Room, fromUser* User, msg string){
	for _, user := range room.Users {
		if fromUser != nil && user.Name == fromUser.Name { // ignore sender itself
			continue
		}
		user.Send(websocket.TextMessage, []byte(msg))
	}
}

func (mgr *RoomMgr) broadCastMsgUserEnter(room *Room, user *User) {
	msg := fmt.Sprintf("{\"type\": \"user_enter\", \"user_name\": \"%s\" }", user.Name)
	mgr.broadCastMsg(room, user, msg)
}

func (mgr *RoomMgr) broadCastMsgUserExit(userName string) {
	room := mgr.getRoomByUserName(userName)
	if room == nil {
		return
	}

	msg := fmt.Sprintf("{\"type\": \"user_exit\", \"user_name\": \"%s\" }", userName)
	mgr.broadCastMsg(room, &User{Name: userName}, msg)
}

func (mgr *RoomMgr) broadCastCustomerMsg(userName string, msg string){
	room := mgr.getRoomByUserName(userName)
	if room == nil {
		return
	}

	mgr.broadCastMsg(room, &User{Name: userName}, msg)
}

func (mgr *RoomMgr) sendRoomUsers2User(room *Room, user *User) {
	if room == nil || user == nil || user.Conn == nil {
		return
	}

	for _, u := range room.Users {
		if u.Name != user.Name {
			msg := fmt.Sprintf("{\"type\": \"user_in_room\", \"user_name\": \"%s\" }", u.Name)
			user.Send(websocket.TextMessage, []byte(msg))
		}
	}
}

func (mgr *RoomMgr) createRoom(roomID string) *Room {
	mgr.mtx.Lock()
	defer mgr.mtx.Unlock()

	r := &Room{
		ID: roomID,
		Users: make(map[string]*User, 0),
	}

	mgr.rooms[roomID] = r
	return r
}

func (mgr *RoomMgr) addUser2Room(room *Room, u *User){
	mgr.mtx.Lock()
	defer mgr.mtx.Unlock()
	room.Users[u.Name] = u
}

func (mgr *RoomMgr) removeUserFromRoom(userName string) {
	mgr.mtx.Lock()
	defer mgr.mtx.Unlock()
	r := mgr.getRoomByUserName(userName)
	if r == nil {
		return
	}

	delete(r.Users, userName)
}

func (mgr *RoomMgr) getRoomByUserName(userName string) *Room {
	for _, room := range mgr.rooms {
		_, ok := room.Users[userName]
		if ok {
			return room
		}
	}
	return nil
}

func (mgr *RoomMgr) getRoomIDByUserName(userName string) string {
	r := mgr.getRoomByUserName(userName)
	if r == nil {
		return ""
	}
	return r.ID
}
func (mgr *RoomMgr) showRoomInfo(roomID string) {
	room, ok := mgr.rooms[roomID]
	if !ok {
		fmt.Printf("Room %s not exists\n", roomID)
		return
	}

	fmt.Printf("====RoomID:%v, UserCnt:%v====\n", roomID, len(room.Users))
	for name, _ := range room.Users {
		fmt.Printf("User: %v\n", name)
	}
	fmt.Printf("==============\n")
}

func (mgr *RoomMgr) UserEnterRoom(user *User, roomID string){
	fmt.Printf("UserEnterRoom: %v, roomID:%v\n", user.Name, roomID)

	// kick out from current room if exists
	mgr.UserExitRoom(user.Name)

	// enter new room
	room, ok := mgr.rooms[roomID]
	if !ok {
		room = mgr.createRoom(roomID)
	}

	// add user to room
	mgr.addUser2Room(room, user)

	// send all other users info to the new user
	mgr.sendRoomUsers2User(room, user)

	// notify user enter to all users in room
	mgr.broadCastMsgUserEnter(room, user)

	// show room info
	mgr.showRoomInfo(roomID)

	// handle connection
	mgr.handleConn(user)
}

func (mgr *RoomMgr) handleConn(user *User){
	fmt.Printf("[HandleConn] name:%v, conn:%v \n", user.Name, user.Conn.RemoteAddr().String())
	// Read data
	for {
		msgType, msg, err := user.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("conn read message failed, name:%v, err:%v \n", user.Name, err)

			// if disconnect by client
			mgr.UserExitRoom(user.Name)
			return
		}

		fmt.Printf("recieve message from %v, msgType:%v, msg:%v \n", user.Name, msgType, string(msg))
		mgr.broadCastCustomerMsg(user.Name, string(msg))
	}
}

func (mgr *RoomMgr) UserExitRoom(userName string){
	roomID := mgr.getRoomIDByUserName(userName)
	if len(roomID) == 0  {
		return
	}
	fmt.Printf("UserExitRoom: %v\n", userName)
	mgr.broadCastMsgUserExit(userName)
	mgr.removeUserFromRoom(userName)
	mgr.showRoomInfo(roomID)
}

