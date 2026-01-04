package main
import "fmt"
import "unicode/utf8"

func  main(){
	var intNum int =  48535
	fmt.Println(intNum)

	var floatNum float64 = 345430.454
	fmt.Println(floatNum)

	var intNum1 int = 6 
	var intNum2 int = 3
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)


	var myString string  = "Hello " + " " + "World"
	fmt.Println(myString)

	fmt.Println(len("Y")) // 2
	fmt.Println(utf8.RuneCountInString("Y")) // 1

	var myRune rune = 'a'
	fmt.Println(myRune) // 97
	

	var myBoolean bool = true
	fmt.Println(myBoolean) // true


	var1 , var2 := 1,2
	fmt.Println(var1,var2)

	const  myConst string = "const value"
	fmt.Println(myConst) 

	
}