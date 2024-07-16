package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("your balance is less than inputed amount")

func CreateAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}

	return &account
}

func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errNoMoney
	}

	account.balance -= amount

	return nil
}

func (account *Account) Deposit(amount int) {
	account.balance += amount
}

func (account *Account) GetBalance() int {
	return account.balance
}

func (account *Account) GetOwner() string {
	return account.owner
}

func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

func (a Account) String() string {
	return fmt.Sprintln(a.GetOwner(), "'s account.\nHas: ", a.GetBalance())
}