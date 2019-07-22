package pointers

import "testing"

func TestWallet(t *testing.T) {
<<<<<<< HEAD
	wallet := Wallet{}

	wallet.Deposit(10)
=======
>>>>>>> 25948db0b935da7094720f591b721ee953cce15c

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(0)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, ErrInsufficientFunds)
	})

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(20)

		assertBalance(t, wallet, want)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
