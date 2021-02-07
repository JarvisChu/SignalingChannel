package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jarvischu/signalchannel/account"
)

// Ping
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//  GetUserList
func GetAccountList(c *gin.Context) {
	rsp := GetAccountListRsp{}
	rsp.AccountList = account.GetAccountList()

	c.JSON(200, &rsp)
}
