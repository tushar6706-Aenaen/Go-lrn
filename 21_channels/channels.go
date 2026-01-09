package main

import (
	"fmt"
	"time"
)

//sending
// func  processNum (numChan chan int){

// 	for num := range numChan {
// 		fmt.Printf("processing number  %d from channel \n", num )
// 		time.Sleep(time.Second)
// 	}
// }

//receive
// func sumNum(result chan int, num1 int, num2 int) {
// 	sumResult := num1 + num2
// 	result <- <-sumResult
// }

// goroutine synchronizer
// func task(done chan bool) {

// 	defer func() { done <- true }()
// 	fmt.Println("Processing...")

// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}
func main() {

	chan1  := make(chan int) 
	chan2  := make(chan string) 


	go func(){
		chan1 <- 10
	}()

	
	go func(){
		chan2 <- "golang"
	}()

	for i:= 0 ; i< 2 ; i++{
		select{
		case chan1Val := <- chan1:
			fmt.Println("received data from chan1",chan1Val)
		case chan2Val := <- chan2:
			fmt.Println("received data from chan2",chan2Val)	
		}
	}
	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan,done)
	// for i := range 100 {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)

	// }

	// fmt.Println("done sending")
	// emailChan <- "tushar@gmail.com"
	// fmt.Println(<-emailChan)

	//important
	// close(emailChan)
	// <-done // block

	// done := make(chan bool)

	// go task(done)

	// <-done // block

	// result := make(chan int)

	// go 	sumNum(result , 9,8)

	// res := <- result // blocking

	// fmt.Println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 67

	// for{
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)
	// messageChan <- "aenaen"

	// msg := <-messageChan
	// fmt.Println(msg)
}
