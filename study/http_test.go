package study

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func login(username, password string) {
	var client = &http.Client{}
	url := "http://106.55.186.222:8200/user/sign-in"
	body := make(map[string]string)
	body["passport"] = username
	body["password"] = password
	bodyByte, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	// 当使用Add时候，如果原本不存在，则添加，如果已存在，就不做任何修改
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	// 当使用Set时候，如果原来这一项已存在，则修改已有的。
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(" http request err:", err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("body close err:", err.Error())
			return
		}
	}(response.Body)

	if response.StatusCode != 200 {
		fmt.Println("http status code err:", err.Error())
		return
	}
	readAll, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read body err:", err.Error())
		return
	}
	fmt.Println(string(readAll))
}

func register(username, password, nickname string) {
	var client = &http.Client{}
	url := "http://106.55.186.222:8200/user/sign-up"
	body := make(map[string]string)
	body["passport"] = username
	body["password"] = password
	body["password2"] = password
	body["nickname"] = nickname
	bodyByte, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	// 当使用Add时候，如果原本不存在，则添加，如果已存在，就不做任何修改
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	// 当使用Set时候，如果原来这一项已存在，则修改已有的。
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(" http request err:", err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("body close err:", err.Error())
			return
		}
	}(response.Body)

	if response.StatusCode != 200 {
		fmt.Println("http status code err:", err.Error())
		return
	}
	readAll, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read body err:", err.Error())
		return
	}
	fmt.Println(string(readAll))
}

func TestLogin(t *testing.T) {
	login("zhangsan", "123456789")
	register("zhangsan", "123456789", "")
}
