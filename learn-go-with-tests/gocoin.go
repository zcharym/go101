package learning_tests

import (
	"fmt"
)

type Dogecoin float64

type Wallet struct {
	balance Dogecoin
}

func (w *Wallet) Deposit(m Dogecoin) {
	fmt.Printf("address:%p\n", &w)
	w.balance += m
}

func (w *Wallet) Balance() Dogecoin {
	fmt.Printf("address:%p\n", &w)
	return w.balance
}
