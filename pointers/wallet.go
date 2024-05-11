package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type Rupee int

func (r Rupee) String() string {
	return fmt.Sprintf("%d Rupee", r)
}

type Wallet struct {
	balance Rupee
}

func (w *Wallet) Deposit(amount Rupee) {
	w.balance += amount
}

func (w *Wallet) Balance() Rupee {
	return w.balance
}

func (w *Wallet) Withdraw(amount Rupee) error {

	if w.balance >= amount {
		w.balance -= amount
		return nil
	}
	return ErrInsufficientFunds
}
