package learning_tests

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{balance: 10}
	wallet.Deposit(10.0)
	got := wallet.Balance()
	expected := Dogecoin(20)

	if got != expected {
		t.Errorf("got '%.2f' expect '%.2f'", got, expected)
	}

}
