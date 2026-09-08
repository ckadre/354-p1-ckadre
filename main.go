package main

import (
	"fmt"

	"github.com/ckadre/354-p1-ckadre/account"
	// "github.com/ckadre/354-p1-ckadre/bank"
	"github.com/ckadre/354-p1-ckadre/customer"
)

func main() {
	var cust = customer.NewCustomer("megumi")
	// var bnk = bank.NewBank(5)
	var savac = account.NewSavingAccount(5, 5.0, cust)
	var chkac = account.NewCheckingAccount(5, 5.0, cust)

	fmt.Println(customer.ToString(cust))

	fmt.Println(chkac.Stringer())

	fmt.Println(savac.Stringer())

	//bnk.add(savac)

	var t int = 0
	fmt.Println(t)
	t += 3
	fmt.Println(t)

}
