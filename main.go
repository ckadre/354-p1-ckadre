package main

import (
	"fmt"

	"github.com/ckadre/354-p1-ckadre/account"
	"github.com/ckadre/354-p1-ckadre/bank"
	"github.com/ckadre/354-p1-ckadre/customer"
)

func main() {
	var cust = customer.NewCustomer("megumi")
	var bnk = bank.NewBank(5)
	var savac = account.NewSavingAccount(5, 5.0, cust)
	var chkac = account.NewCheckingAccount(5, 5.0, cust)

	fmt.Println(customer.ToString(cust))

	fmt.Println(chkac.ToString())

	chkac.Deposit(100.0)

	fmt.Println(savac.ToString())

	fmt.Println(bnk.ToString())

	bnk.Add(savac)
	bnk.Add(chkac)

	fmt.Println(bnk.ToString())
}
