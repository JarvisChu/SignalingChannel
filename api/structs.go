package api

import "github.com/jarvischu/signalchannel/account"

type GetAccountListRsp struct {
	Code        int                `json:"code"`
	Message     string             `json:"msg"`
	AccountList []*account.Account `json:"accounts"`
}
