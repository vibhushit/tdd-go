package wallet

import (
	"errors"
	"fmt"
)

var insufficientFundsError = "can not withdraw, insufficient funds"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(val Bitcoin) {
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	w.balance += val
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New(insufficientFundsError)
	}

	w.balance -= amount
	return nil
}
