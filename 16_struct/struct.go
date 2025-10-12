package main

import (
	"fmt"
	"time"
)

type order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time // nanosecond
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup
	myOrder := order{
		id:       id,
		amount:   amount,
		status:   status,
		createAt: time.Now(),
	}

	return &myOrder
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func main() {
	myOrder := order{
		id:     "234",
		amount: 234.34,
		status: "pending",
	}

	myOrder.createAt = time.Now()

	fmt.Println(myOrder.amount)

	fmt.Println("My Order", myOrder)

	myOrder2 := order{
		id:       "2",
		status:   "delivery",
		amount:   324.32,
		createAt: time.Now(),
	}

	myOrder.status = "paid"

	myOrder2.changeStatus("confirm")

	fmt.Println(myOrder, myOrder2)

	fmt.Println(myOrder.getAmount())

	myOrderX := newOrder("1", 123.32, "pending")

	fmt.Println(myOrderX.amount)

	language := struct {
		name   string
		isGood bool
	}{"Go Lang", true}

	fmt.Println(language)
}
