package customer

type Customer struct {
	name string //name of the customer
}

func NewCustomer(name string) (c *Customer) {
	c = new(Customer)
	c.name = name
	return c
}

func ToString(c *Customer) (name string) {
	return c.name
}
