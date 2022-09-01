package main

import (
	"encoding/json"
	"fmt"
)

type proper struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height float32
}

func main() {
	stu_map := make(map[string]int)
	stu_map["zhangsan"] = 8
	stu_map["lisi"] = 8
	stu_map["wangwu"] = 8
	fmt.Println(stu_map)
	//编码成json
	res, err := json.MarshalIndent(stu_map, "", " ")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("res:", string(res))

	//json转成map
	abc := make(map[string]int)
	err2 := json.Unmarshal([]byte(res), &abc)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	fmt.Println("json:", abc)

	//对象转json
	p := proper{Name: "张三", Age: 15, Height: 19.8}
	fmt.Println(p)
	res2, err3 := json.MarshalIndent(p, "", " ")
	if err3 != nil {
		fmt.Println("error", err3)
		return
	}
	fmt.Println("res:", string(res2))
}
