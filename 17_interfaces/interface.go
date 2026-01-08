package main

import "fmt"

type paymenter interface{
	pay(amount float32) 
	refund(amount float32,account string)
}



type payment struct {
	gateway paymenter

}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw :=stripe{}
	// stripePaymentGw.pay(amount)
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay ",amount)
}

type stripe struct {}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe ",amount)

}

type fake struct {}

func (f fake) pay(amount float32){
	fmt.Println("making payment using fakePayment ",amount)

}
func (f fake) refund(amount float32,account string){
	fmt.Println("making refund using fakePayment ",amount,account)
}
func main() {

	fakePaymentGw := fake{}
	// stripePaymentGw := stripe{}
	newPaymet := payment{
		gateway: fakePaymentGw,
	}
	newPaymet.makePayment(100)
}