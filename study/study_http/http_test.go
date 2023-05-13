package study_http

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	// 调用封装的接口请求
	login("zhangsan", "123456789")
	register("zhangsan", "123456789", "")
	fmt.Println("======================")

	// 看一下封装的请求是否好用
	response, err := httpPost("http://106.55.186.222:8200/user/sign-in", "{Passport}")
	if err != nil {
		fmt.Println("http post err:", err.Error())
		return
	}
	fmt.Println(string(response))

}
