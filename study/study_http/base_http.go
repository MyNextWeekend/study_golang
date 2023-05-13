package study_http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// 封住一个post请求方法
func httpPost(url string, body string) (result []byte, err error) {
	client := &http.Client{}

	request, _ := http.NewRequest("post", url, strings.NewReader(body))
	// 当使用Add时候，如果原本不存在，则添加，如果已存在，就不做任何修改
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	// 当使用Set时候，如果原来这一项已存在，则修改已有的。
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(" http request err:", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("http status code err:" + strconv.Itoa(response.StatusCode))
	}
	result, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("read body err:" + err.Error())
	}
	return result, nil
}
