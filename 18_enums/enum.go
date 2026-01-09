package main

import "fmt"

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderderStatus(status OrderStatus) {
	fmt.Println(status)
}

func main() {
	changeOrderderStatus(Confirmed)
}
