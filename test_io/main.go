package main

import (
	"fmt"
	"os"
)

func main() {
	readFile, err := os.ReadFile("./test_io/main.go")
	if err == nil {
		fmt.Println("文件的长度是：", len(readFile))
		fmt.Println(len(string(readFile)))
		fmt.Println(string(readFile[0:150]))
		//for i, b := range readFile {
		//	fmt.Println("index:", i, "bin:", b)
		//}
	} else {
		fmt.Print("异常了呀：", err)
	}
}
