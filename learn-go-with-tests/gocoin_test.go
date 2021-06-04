package learning_tests

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{balance: Dogecoin(10)}
	wallet.Deposit(10.0)
	wallet.Withdraw(4.0)
	got := wallet.Balance()
	expected := Dogecoin(16)

	if got != expected {
		t.Errorf("got '%.3f' expect '%.3f'", got, expected)
	}

}
