package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/signalchannel/account"
	"github.com/jarvischu/signalchannel/ws"
	"net/http"
)

// Ping
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetUserList
func GetAccountList(c *gin.Context) {
	rsp := GetAccountListRsp{}
	rsp.AccountList = account.GetAccountList()

	c.JSON(http.StatusOK, &rsp)
}

// SendMsg
func SendMsg(c *gin.Context) {
	req := SendMsgReq{}
	rsp := SendMsgRsp{}

	// parse param
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// check
	a, err := account.GetAccountByID(req.Recipient)
	if err != nil {
		rsp.Code, rsp.Message = -1, err.Error()
		c.JSON(http.StatusOK, &rsp)
		return
	}

	if a.Status == account.Offline {
		rsp.Code, rsp.Message = -1, "recipient is offline"
		c.JSON(http.StatusOK, &rsp)
		return
	}

	// send msg
	if err := ws.GetConnMgr().Send(req.Sender, req.Recipient, req.Message); err != nil {
		rsp.Code, rsp.Message = -1, "send msg failed"
		c.JSON(http.StatusOK, &rsp)
		return
	}

	c.JSON(http.StatusOK, &rsp)
}
