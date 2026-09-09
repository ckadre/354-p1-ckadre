package account

import (
	"strconv"

	"github.com/ckadre/354-p1-ckadre/customer"
)

type checkingAccount struct {
	number int
	bal    float32
	cust   *customer.Customer
}

func NewCheckingAccount(num int, bal float32, cust *customer.Customer) (ca *checkingAccount) {
	ca = new(checkingAccount)
	ca.number = num
	ca.bal = bal
	ca.cust = cust
	return ca
}

func (ca *checkingAccount) accrue(rate float32) {
	//checking accounts don't accrue interest
}

func (ca *checkingAccount) balance() float32 {
	return ca.bal
}

func (ca *checkingAccount) deposit(amount float32) {
	ca.bal += amount
}

func (ca *checkingAccount) withdraw(amount float32) {
	ca.bal -= amount
}

func (ca *checkingAccount) ToString() (rtn string) {
	return strconv.Itoa(ca.number) + ": " + customer.ToString(ca.cust) + ": " + strconv.FormatFloat(float64(ca.bal), 'f', -1, 64)
}
