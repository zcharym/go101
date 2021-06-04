package learning_tests

import (
	"fmt"
)

type Dogecoin float64

type Wallet struct {
	balance Dogecoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(m Dogecoin) {
	w.balance += m
}

func (w *Wallet) Withdraw(m Dogecoin) {
	w.balance -= m
}

func (w *Wallet) Balance() Dogecoin {
	return w.balance
}

func (d Dogecoin) String() string {
	return fmt.Sprintf("%.3f DOGE", d)
}
