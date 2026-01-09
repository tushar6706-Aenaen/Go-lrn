package main

import (
	"fmt"
	"sync"
)

func printnums(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go printnums(i, &wg)
	}

	wg.Wait()
}
