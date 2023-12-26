package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		if wallet.balance != want {
			t.Errorf("got %s want %s", wallet.balance, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted and error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))

		// got := wallet.Balance()

		// want := Bitcoin(10)

		// fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(15))

		// got := wallet.Balance()
		// want := Bitcoin(15)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})

	t.Run("insufficient withdrawal error", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, insufficientFundsError)

		assertBalance(t, wallet, initialBalance)

		// if err == nil {
		// 	t.Errorf("should give an error but didn't")
		// }

	})
}
