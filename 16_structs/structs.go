package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func newOrder(id string, amount float32, status string) *order {

	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder

}

// revceiver
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func main() {
	// newCustomer :=customer{
	// 	name:"tung",
	// 	phone:"7566555435",
	// }
	newOrder := order{
		id:     "101",
		amount: 302,
		status: "paid",
		customer: customer{
			name:  "tung",
			phone: "7566555435",
		},
	}

	fmt.Println(newOrder)
	// language := struct{
	// 	name string
	// 	isGood bool
	// } {"golang",true}

	// fmt.Println(language)

	// myOrder := newOrder("1",50.2,"received")
	// fmt.Println(myOrder)

	// myOrder := order{
	// id:     "1",
	// amount: 50.00,
	// status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder.getAmount())

	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2:=order{
	// 	id:"1",
	// 	amount:50.00,
	// 	status: "recieved",

	// }

	// myOrder.status = "paid"

	// fmt.Println("Order struct ",myOrder)
	// fmt.Println("Order struct ",myOrder2)

}
