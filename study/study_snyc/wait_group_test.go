package study_snyc

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func hello(num int) {
	//计数器减一
	defer wg.Done()
	fmt.Println("hello", num)
}

func TestSnyc(t *testing.T) {
	//计数器增加计数
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	//等待计数器归零
	wg.Wait()
}
