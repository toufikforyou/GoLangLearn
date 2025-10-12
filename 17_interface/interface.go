package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	getway paymenter
}

func (p payment) makePayment(amount float32) {
	// stripePay := stripe{}
	// stripePay.pay(amount)
	p.getway.pay(amount)
}

type bkash struct{}

func (r bkash) pay(amount float32) {
	// login for payment pay
	fmt.Println("Making payment using bkash pay", amount)
}

type stripe struct{}

func (r stripe) pay(amount float32) {
	// login for payment pay
	fmt.Println("Making payment using stripe pay", amount)
}

func main() {
	newPayment := payment{
		getway: stripe{},
	}

	newPayment.makePayment(23)

}
