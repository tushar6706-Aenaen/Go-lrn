package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderderStatus(status OrderStatus) {
	fmt.Println(status)
}

func main() {
	changeOrderderStatus(Confirmed)
}