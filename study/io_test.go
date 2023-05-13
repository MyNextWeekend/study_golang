package study

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func FileToBytes(filePath string, size int) (res [][]byte) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print("Open File Error:", err.Error())
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

func test02() (res [][]byte) {
	readFile, err := os.Open("./io_test.go")
	if err != nil {
		fmt.Print("Open File Error:", err.Error())
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

func TestIO(t *testing.T) {
	//bytes := FileToBytes("./test_io/main.go", 3200)
	//fmt.Println(bytes)
	test02()
}
