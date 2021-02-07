package account

import (
	"fmt"
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
	Calling Status = "calling"
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

func GetAccountByID(id string)(*Account, error){
	for _, a := range accounts {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, fmt.Errorf("not fount")
}