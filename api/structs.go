package api

import "github.com/jarvischu/signalingchannel/account"

type GetAccountListRsp struct {
	Code        int                `json:"code"`
	Message     string             `json:"msg"`
	AccountList []*account.Account `json:"accounts"`
}

type SendMsgReq struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Message   string `json:"msg"`
}

type SendMsgRsp struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
