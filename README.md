# SignalChannel

SignalChannel for WebRTC

## How to build and run

```
go build .
./signalchannel
```

## How to test

1. build SignalChannel
2. run SignalChannel
3. open and using [WebSocket Echo Page](https://www.websocket.org/echo.html) with Location `ws://localhost:8080/ws?id=1` to connect
4. repeat step 3, with `id=2`, so you got two users (1,2) connected to SignalChannel
5. call HTTP API `/msg` to send msg each other, e.g.

```json
POST localhost:8080/msg

{
    "sender":"1",
    "recipient":"2",
    "msg": "hello"
}
```