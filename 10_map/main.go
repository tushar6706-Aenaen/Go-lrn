package main

import (
	"fmt"
	"maps"
)

func main() {

	//creating map

	// m := make(map[string]string)
	// setting an element
	// m["language"] = "golang"
	// m["area"] = "backend"

	//get an element
	// fmt.Println(m["language"],m["area"])
	//imp : if key does not exist  then map return zero values
	// fmt.Println(m["someone"])

	// m := make(map[string]int)
	// m["area"] = 20
	// m["amount"] = 60
	// fmt.Println(m["someone"])

	// fmt.Println(len(m))

	// delete(m, "area")
	// fmt.Println(m)
	// clear(m)
	// fmt.Println(m)

	// m := map[string]string{"laptop": "hp pavillion"}
	// fmt.Println(m)

	// value, ok := m["laptop"]
	// fmt.Println(value)

	// if ok {
	// 	fmt.Print("all Ok")
	// } else {
	// 	fmt.Print("not Ok")

	// }

	m1 := map[string]string{"laptop": "hp pavillion"}
	m2 := map[string]string{"laptop": "hp pavillion"}

	fmt.Println(maps.Equal(m1, m2))
}
