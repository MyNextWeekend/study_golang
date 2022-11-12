package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU核数：", runtime.NumCPU())
	go func(s string) {
		defer fmt.Println("结束")
		for i := 0; i < 2; i++ {
			runtime.Goexit()
			fmt.Println(s)
			runtime.Gosched() // 让出CPU时间片，重新等带安排任务
		}
		defer fmt.Println("结束2")
	}("world")

	for i := 0; i < 2; i++ {
		fmt.Println("hello")
		runtime.Gosched() // 让出CPU时间片，重新等待安排任务
	}
}
