package account

import (
	"github.com/ckadre/354-p1-ckadre/customer"
)

type savingAccount struct {
	number int
	bal    float32
	cust   customer.Customer
}

func newSavingAccount(num int, bal float32, cust customer.Customer) (sa *savingAccount) {
	sa = new(savingAccount)
	sa.bal = bal
	sa.number = num
	sa.cust = cust
	return sa
}

func (sa *savingAccount) accrue(rate float32) {
	sa.bal *= rate
}

func (sa *savingAccount) balance() (bal float32) {
	return sa.bal
}

func (sa *savingAccount) deposit(amount float32) {
	sa.bal += amount
}

func (sa *savingAccount) withdraw(amount float32) {
	sa.bal -= amount
}
