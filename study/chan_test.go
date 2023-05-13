package study

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {

	var c = make(chan int)

	go func() {
		//写在方法的第一行记住关闭通道
		defer close(c)
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()

	//已关闭的通道是可以循环遍历的
	for i := range c {
		fmt.Println("收到的数据是：", i)
	}
}
