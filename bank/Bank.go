package bank

import (
	"math"
	"strconv"

	"github.com/ckadre/354-p1-ckadre/account"
)

type Bank struct {
	accounts []account.Accounter
	i        int
	interest float32
}

func NewBank(len int) (b *Bank) {
	b = new(Bank)
	b.accounts = make([]account.Accounter, len)
	b.i = 0
	b.interest = 0.0
	return b
}

func (b *Bank) Add(Account account.Accounter) {
	b.accounts[b.i] = Account
	b.i += 1
}

func (b *Bank) Accrue(rate float32) {
	for v := range b.i {
		b.interest += account.Accounter.Accrue(b.accounts[v], rate)
	}
}

func (b *Bank) ToString() (bnk string) {
	var str string
	for v := range b.i {
		str += account.Accounter.ToString(b.accounts[v])
		str += "\n"
	}
	str += "Total Interest: " + strconv.FormatFloat(math.Round(float64(b.interest)), 'f', 2, 64)
	return str
}
