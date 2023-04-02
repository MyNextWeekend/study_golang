package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type proper struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height float32
}

func mapToJson() []byte {
	stuMap := make(map[string]int)
	stuMap["zhangsan"] = 8
	stuMap["lisi"] = 9
	stuMap["wangwu"] = 10
	fmt.Println("map:", stuMap)
	//编码成json
	res, err := json.MarshalIndent(stuMap, "", " ")
	if err != nil {
		fmt.Println("MarshalIndent Error:", err.Error())
		return nil
	}
	fmt.Println("json_str:", string(res))
	return res
}

// json转成map
func jsonToMap(jsonByte []byte) {
	resMap := make(map[string]int)
	err := json.Unmarshal(jsonByte, &resMap)
	if err != nil {
		fmt.Println("Unmarshal Error:", err.Error())
	}
	fmt.Println("map:", resMap)
}

// 对象转json
func objToJson() {
	zhangSan := proper{Name: "张三", Age: 15, Height: 19.8}
	fmt.Println(zhangSan)
	res2, err := json.MarshalIndent(zhangSan, "", " ")
	if err != nil {
		fmt.Println("MarshalIndent Error:", err.Error())
		return
	}
	fmt.Println("json_str:", string(res2))
}
func TestMap(t *testing.T) {
	//objToJson()
	jsonByte := mapToJson()
	jsonToMap(jsonByte)
}
