package pointers

import (
	"errors"
	"fmt"
)

func main() {

}

type Bitcoin int

var ErrInsufficientFunds = errors.New("Attempted to overdraw")

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

<<<<<<< HEAD
func (w Wallet) Balance() int {
	return 0
=======
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
>>>>>>> 25948db0b935da7094720f591b721ee953cce15c
}
