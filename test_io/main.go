package main

import (
	"fmt"
	"os"
)

func FileToBytes(filePath string, size int) (res [][]byte) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < (len(file)/size)+1; i++ {
		start := i * size
		end := (i + 1) * size
		if len(file) < end {
			end = len(file)
		}
		res = append(res, file[start:end])
	}
	return
}

func test() {
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

func main() {
	bytes := FileToBytes("./test_io/main.go", 3200)
	fmt.Println(bytes)
}
