# SignalingChannel

SignalingChannel for WebRTC

## How to build and run

```bash
make
./signalingchannel
```

or

```bash
./run.sh
```

## How to test

1. build SignalingChannel
2. run SignalingChannel
3. open and using [WebSocket Echo Page](https://www.websocket.org/echo.html) with Location `ws://localhost:8080/ws?id=1` to connect
4. repeat step 3, with `id=2`, so you got two users (1,2) connected to SignalingChannel
5. call HTTP API `/msg` to send msg each other, e.g.

```json
POST localhost:8080/msg

{
    "sender":"1",
    "recipient":"2",
    "msg": "hello"
}
```