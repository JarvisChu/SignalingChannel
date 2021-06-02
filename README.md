# SignalingChannel

SignalingChannel for WebRTC, support P2P mode and Room mode

**P2P mode**
Signalling for Peer 2 Peer mode, and Two Peers can signalling directly.

**Room mode**
Using for meeting, assign room for Peers, and Peers in the same room can communicating with each other.


## How to build and run

```bash
./run.sh
```

## How to use

### P2P mode
1. Run SignalingChannel by `run.sh`
2. User1 Connect to SignalChannel by websocket with URL `ws://localhost:8080/ws/p2p?name=User1`
3. User2 Connect to SignalChannel by websocket with URL `ws://localhost:8080/ws/p2p?name=User2`
4. User1 `websocket.send("set-peer:User2")`
5. Then, any data send through websocket by User1, will directly send to User2
6. User2 `websocket.send("set-peer:User1")`
7. Then, any data send through websocket by User2, will directly send to User1

See webrtc client demo: https://github.com/JarvisChu/WebRTCClient  p2p

### Room mode
1. Run SignalingChannel by `run.sh`
2. User1 Connect to SignalChannel by websocket with URL `ws://localhost:8080/ws/room?name=User1&roomid=12345`
3. User2 Connect to SignalChannel by websocket with URL `ws://localhost:8080/ws/p2p?name=User2&roomid=12345`
4. UserN Connect to SignalChannel by websocket with URL `ws://localhost:8080/ws/p2p?name=UserN&roomid=12345`
5. Then, User1,User2,..,UserN are in the same room (roomid: 12345). Any data send through websocket by any user will boardcast to all other users in the room.

See webrtc client demo: https://github.com/JarvisChu/WebRTCClient  room
