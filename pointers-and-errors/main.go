package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}
type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	w.balance += amount
	return nil
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
