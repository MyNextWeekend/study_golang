package main

import (
	"encoding/json"
	"fmt"
)

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

}
