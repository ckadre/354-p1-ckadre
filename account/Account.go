package account

type Accounter interface {
	balance() float32

	accrue(rate float32)

	deposit(amount float32)

	withdraw(amount float32)

	ToString() (rtn string)
}
