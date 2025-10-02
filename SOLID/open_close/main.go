package main

/*
	We can add new payment method (crad payment)
	without disturbing other payment methods
*/

import "fmt"

type PaymentProcessor interface {
	Process()
}

type CashMethod struct{}

func (c *CashMethod) Process() {
	fmt.Println("Processing with cash")
}

type UPI struct{}

func (u *UPI) Process() {
	fmt.Println("Processing with UPI")
}

func HandleProcess(p PaymentProcessor) {
	p.Process()
}

func main() {
	c := CashMethod{}
	HandleProcess(&c)
	// both line 26+27 and 29 are same
	HandleProcess(&UPI{})
}
