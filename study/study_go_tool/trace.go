package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/*
执行过程可视化
1、执行go run trace.go
2、生成trace.out
3、执行 go tool trace trace.out
*/
func main() {
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	err = trace.Start(file)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("hello world...")

	trace.Stop()
}
