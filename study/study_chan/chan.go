package main

import (
	"fmt"
)

func one() {

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

func two() {
	var c = make(chan string)

	go func() {
		defer close(c)
		for {
			var value string
			fmt.Println("请输入一个值")
			fmt.Scanln(&value)
			if value == "q" {
				return
			}
			c <- value
		}
	}()

	for i := range c { // 程序在此处会阻塞，等待通道中所有任务消费完成（即：通道不关闭，永远阻塞）
		//time.Sleep(5 * time.Second)
		fmt.Println("收到数据：", i)
	}
	println("done")
}

func main() {
	two()
}
