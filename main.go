package main

import (
	"fmt"

	"github.com/ckadre/354-p1-ckadre/account"
	"github.com/ckadre/354-p1-ckadre/bank"
	"github.com/ckadre/354-p1-ckadre/customer"
)

func main() {
	var bnk = bank.NewBank(5) //functionality could easily be improved by not hard coding the length of the bank slice. ¯\_(ツ)_/¯
	var ann = customer.NewCustomer("Ann")
	var bob = customer.NewCustomer("Bob")

	bnk.Add(account.NewCheckingAccount(1, 100.00, ann))
	bnk.Add(account.NewSavingAccount(2, 200.00, ann))
	bnk.Add(account.NewSavingAccount(3, 150.00, bob))

	bnk.Accrue(0.02)

	fmt.Println(bnk.ToString())
}
