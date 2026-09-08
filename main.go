package main

import (
	"fmt"
	// "github.com/ckadre/354-p1-ckadre/account"
	// "github.com/ckadre/354-p1-ckadre/bank"
	"github.com/ckadre/354-p1-ckadre/customer"
)

func main() {
	var cust = customer.NewCustomer("megumi")
	// var b = bank.newBank(5)
	// var savac = account.newSavingAccount(5, 5, cust)
	fmt.Println(customer.ToString(cust))

	var t int = 0
	fmt.Println(t)
	t += 3
	fmt.Println(t)

}
