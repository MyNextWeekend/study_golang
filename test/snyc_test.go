package test

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func hello(num int) {
	defer wg.Done()
	fmt.Println("hello", num)
}

func TestSnyc(t *testing.T) {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	wg.Wait()
}
