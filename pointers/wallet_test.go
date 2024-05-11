package pointers

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Rupee) {
		t.Helper()
		got := wallet.Balance()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if !(errors.Is(want, got)) {
			t.Errorf("want %s got %s", want, got)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		tenRupee := Rupee(10)
		wallet.Deposit(tenRupee)

		assertBalance(t, wallet, Rupee(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		tenRupee := Rupee(10)
		fiveRupee := Rupee(5)
		wallet.Deposit(tenRupee)
		err := wallet.Withdraw(fiveRupee)

		assertBalance(t, wallet, Rupee(5))
		if err != nil {
			return
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		tenRupee := Rupee(10)
		twentyRupee := Rupee(20)
		wallet.Deposit(tenRupee)
		err := wallet.Withdraw(twentyRupee)

		assertError(t, err, ErrInsufficientFunds)
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
