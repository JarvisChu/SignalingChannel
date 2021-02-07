package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func Handle(c *gin.Context) {
	upgrader := websocket.Upgrader{
		// CheckOrigin：请求检查函数，用于统一的链接检查，以防止跨站点请求伪造。如果不检查，就设置一个返回值为true的函数。
		// 如果请求Origin标头可以接受，CheckOrigin将返回true。
		// 如果CheckOrigin为nil，则使用安全默认值：如果Origin请求头存在且原始主机不等于请求主机头，则返回false
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	uid := c.Query("uid")
	if len(uid) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	GetConnMgr().HandleConn(uid, conn)
	return
}
