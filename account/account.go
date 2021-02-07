package account

import (
	"sync"
)

var (
	mtx      sync.Mutex
	accounts []*Account
)

// AccountStatus
type Status string

const (
	Offline Status = "offline"
	Online  Status = "online"
	Calling Status = "calling" // 通话中
)

// Account
type Account struct {
	ID string //
	Status
}

func GetAccountList() []*Account {
	return accounts
}

func AddAccount(account *Account) {
	mtx.Lock()
	defer mtx.Unlock()

	for _, a := range accounts {
		if a.ID == account.ID {
			return
		}
	}

	accounts = append(accounts, account)
}

func UpdateAccountStatus(id string, status Status) {
	mtx.Lock()
	defer mtx.Unlock()

	for _, a := range accounts {
		if a.ID == id {
			a.Status = status
			return
		}
	}

	return
}
