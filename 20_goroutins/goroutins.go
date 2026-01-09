package main

import (
	"fmt"
	"time"
)

func printnums(id int) {
	fmt.Println(id)
}

func main() {

	for i:= 0 ; i<= 10 ; i++{
		go printnums(i)
	}

	time.Sleep(time.Second)

}