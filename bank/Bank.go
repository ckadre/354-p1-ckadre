package bank

import (
	"github.com/ckadre/354-p1-ckadre/account"
)

type Bank struct {
	accounts []account.Accounter
	i        int
}

func NewBank(len int) (b *Bank) {
	b = new(Bank)
	b.accounts = make([]account.Accounter, len)
	b.i = 0
	return b
}

func (b *Bank) Add(Account account.Accounter) {
	b.accounts[b.i] = Account
	b.i += 1
}

func (b *Bank) accrue(rate float32) {

}

func (b *Bank) ToString() (bnk string) {
	var str string = ""
	for v := range b.i {
		str += account.Accounter.ToString(b.accounts[v])
		str += "\n"
	}
	return str
}
