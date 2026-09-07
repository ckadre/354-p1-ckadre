package bank

import (
	"hash"
	"github.com/ckadre/354-p1-ckadre/account"
)

type Bank struct {
	accounts map[*IAccount]IAccount
}

func newBank()(b *Bank) {
	b = new(Bank) 
	b.accounts = Bank [account]account
}

func add(Account account) {
	accounts.add(account)
}