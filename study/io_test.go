package study

import (
	"fmt"
	"os"
	"testing"
)

// 读取文件，按照传入size切割
func FileToBytes(filePath string, size int) (res [][]byte, err error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	for i := 0; i < (len(file)/size)+1; i++ {
		start := i * size
		end := (i + 1) * size
		if len(file) < end {
			end = len(file)
		}
		res = append(res, file[start:end])
	}
	return res, nil
}

func TestFileToBytes(t *testing.T) {
	bytes, err := FileToBytes("./study_generate_dir/file.json", 20)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(len(bytes))
	for _, value := range bytes {
		fmt.Println(string(value))
	}
}
