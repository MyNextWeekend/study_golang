package main

import (
	"fmt"
	"io"
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
	readFile, err := os.Open("./test_io/main.go")
	if err != nil {
		fmt.Print("异常了呀：", err)
	}
	buf := make([]byte, 1024)
	n, err := readFile.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(len(buf))
}

func test02() (res [][]byte) {
	readFile, err := os.Open("./test_io/main.go")
	if err != nil {
		fmt.Print("异常了呀：", err)
	}
	for {
		buf := make([]byte, 10)
		n, err := readFile.Read(buf)
		if err == io.EOF {
			break
		}
		res = append(res, buf[:n])
	}
	return
}

func main() {
	//bytes := FileToBytes("./test_io/main.go", 3200)
	//fmt.Println(bytes)
	test02()
}
