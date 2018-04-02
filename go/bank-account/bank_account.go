package account

import "sync"

type Account struct {
	balance int64
	// m       *sync.Mutex
	rw   *sync.RWMutex
	open bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, &sync.RWMutex{}, true}
}

func (a *Account) Balance() (int64, bool) {
	if !a.open {
		return -1, false
	}

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.rw.Lock()
	defer a.rw.Unlock()
	if a.open {
		a.open = false
		payout := a.balance
		a.balance = 0
		return payout, true
	}

	return a.balance, false
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.open {
		a.rw.Lock()
		defer a.rw.Unlock()
		newBalance := a.balance + amount
		if newBalance >= 0 {
			a.balance = newBalance
			return a.balance, true
		}

		return a.balance, false

	}

	return 0, false
}
