package main

import "fmt"

var Name string

func init() {
	fmt.Printf("初始化之前的数据%v\n", Name)
	Name = "zhangsan"
	fmt.Printf("初始化之后的数据%v\n", Name)
}

func main() {
	fmt.Print("666")
}
