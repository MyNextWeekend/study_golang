package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	wg.Wait()
}

func hello(num int) {
	defer wg.Done()
	fmt.Println("hello", num)
}
