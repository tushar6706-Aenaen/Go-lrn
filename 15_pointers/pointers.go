package main

import "fmt"

func change(num *int) {
	*num = 100
	fmt.Println("in change num",*num)
}

func main() {

	num := 1
	change(&num)

	fmt.Println("after change ",num)

}