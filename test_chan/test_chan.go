package main

import (
	"fmt"
)

func main() {

	var c = make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println("shoudao", i)
	}
}
