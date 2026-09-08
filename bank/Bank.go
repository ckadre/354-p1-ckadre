package bank

import (
	// "hash"
	"github.com/ckadre/354-p1-ckadre/account"
)

type Bank struct {
	accounts []account.Accounter
	i        int
}

func newBank(len int) (b *Bank) {
	b = new(Bank)
	b.accounts = make([]account.Accounter, len)
	b.i = 0
	return b
}

func add(b *Bank, Account account.Accounter) {
	b.accounts[b.i] = Account
	b.i += 1
}

func accrue(b *Bank, rate float32) {

}

func ToString(b *Bank) (bnk string) {
	for  := range b.i {
		bnk.append()
	}
}
