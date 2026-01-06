package main

import "fmt"

// func getLanguages() (string,string,bool){
// 	return "golang" ,"javascript" , true
// }



// func add(a int, b int) int {
// 	return a + b
// }

// func processIt(fn func(a int )int){
// 	fn(1)
// }

func processIt()func(a int )int{
	return func(a int)int{
		return a+60
	}
}
func main() {
	fn:=processIt()
	fn(9)
	fmt.Println(fn(9))
	// fn:=func(a int )int {
	// 	return 2
	// }
	// processIt(fn)	

	// lang1 , lang2,_ := getLanguages()
	// fmt.Println(lang1,lang2)
	// result := add(1, 2)
	// fmt.Println(result)




}