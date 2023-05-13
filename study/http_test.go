package study

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// get请求
func TestDoGet(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("Http Error:", err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Body Error:", err.Error())
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
func TestDoPost(t *testing.T) {
	proper := Proper{Name: "张三", Age: 18}
	jsonByte, err := json.Marshal(&proper)
	if err != nil {
		fmt.Println("Json Marshal Error", err.Error())
		return
	}
	fmt.Println("转换结果：", jsonByte)

	resp, err := http.Post(`https://www.baidu.com`, "application/json", bytes.NewBuffer(jsonByte))
	if err != nil {
		fmt.Println("Http Error:", err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Http Code not 200:")
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Body Error:", err.Error())
		return
	}
	fmt.Println(string(body))

}
