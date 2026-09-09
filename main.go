package main

import (
	"fmt"

	"github.com/ckadre/354-p1-ckadre/account"
	"github.com/ckadre/354-p1-ckadre/bank"
	"github.com/ckadre/354-p1-ckadre/customer"
)

func main() {
	var bnk = bank.NewBank(5)
	var ann = customer.NewCustomer("Ann")
	var bob = customer.NewCustomer("Bob")

	bnk.Add(account.NewCheckingAccount(1, 100.00, ann))
	bnk.Add(account.NewSavingAccount(2, 200.00, ann))
	bnk.Add(account.NewSavingAccount(3, 150.00, bob))

	bnk.Accrue(0.02)

	fmt.Println(bnk.ToString())

	// var savac = account.NewSavingAccount(5, 5.0, cust)
	// var chkac = account.NewCheckingAccount(5, 0.0, cust)

	// fmt.Println(customer.ToString(ann))

	// fmt.Println(chkac.ToString())

	// chkac.Deposit(100.0)

	// fmt.Println(savac.ToString())

	// fmt.Println(bnk.ToString())

	// bnk.Add(savac)
	// bnk.Add(chkac)

	// fmt.Println(bnk.ToString())

	// bnk.Accrue(0.1)

	// fmt.Println(bnk.ToString())

}
