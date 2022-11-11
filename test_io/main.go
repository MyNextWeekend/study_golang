package main

import (
	"fmt"
	"os"
)

func main() {
	readFile, err := os.ReadFile("./test_io/main.go")
	if err == nil {
		fmt.Print("文件的长度是：", len(readFile))
	} else {
		fmt.Print("异常了呀：", err)
	}
}
