package account

type Accounter interface {
	Balance() float32

	Accrue(rate float32) (total float32)

	Deposit(amount float32)

	Withdraw(amount float32)

	ToString() (rtn string)
}
