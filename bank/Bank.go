package bank

import (
	"github.com/ckadre/354-p1-ckadre/account"
)

type Bank struct {
	accounts []account.Account
	i int
}

func NewBank(len int) (b *Bank) {
	b = new(Bank)
	b.accounts = make([]account.Account, len)
	b.i = 0
	return b
}

func (b *Bank) add(Account account.Account) {
	b.accounts[b.i] = Account
	b.i += 1
}

func (b *Bank) accrue(rate float32) {

}

func (b *Bank) ToString() (bnk string) {
	for  := range b.i {
		bnk.append()
	}
}
