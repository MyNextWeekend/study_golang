package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// get请求
func doGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(respBytes))
	fmt.Println(resp.Status)
}

type Proper struct {
	Name string
	Age  int
}

// 结构体转json
func doPost() {
	proper := Proper{Name: "张三", Age: 18}
	marshal, err2 := json.Marshal(&proper)
	if err2 != nil {
		fmt.Println("转化失败。")
		return
	}
	fmt.Println("转换结果：", marshal)
	resp, err := http.Post(`http://www.baidu.com`, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Print(resp.Status)
}

func TestHttp(t *testing.T) {
	doGet()
	doPost()
}
