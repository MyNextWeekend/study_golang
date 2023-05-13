package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var user *User
	user = &User{
		Id:   3,
		Name: "张三",
		Age:  18,
	}

	// 结构体转 字符串/map
	userStr := gconv.String(user)
	userMap := gconv.Map(user)
	fmt.Printf("userStr value:%v  userStr type:%T\n", userStr, userStr)
	fmt.Printf("userMap value:%v  userMap type:%T\n", userMap, userMap)
	fmt.Println("-----------")

	// 字符串转 map/结构体指针
	userMap = gconv.Map(userStr)
	if gconv.Struct(userStr, &user) != nil {
		fmt.Println("string to struct err")
	}
	fmt.Printf("userMap value:%v  userMap type:%T\n", userMap, userMap)
	fmt.Printf("user value:%v  user type:%T\n", user, user)
	fmt.Println("-----------")

	// map转 字符串/结构体指针
	userStr = gconv.String(userMap)
	if gconv.Struct(userMap, &user) != nil {
		fmt.Println("map to struct err")
	}
	fmt.Printf("userStr value:%v  userStr type:%T\n", userStr, userStr)
	fmt.Printf("user value:%v  user type:%T\n", user, user)
}
