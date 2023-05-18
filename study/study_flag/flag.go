package main

import (
	"flag"
	"fmt"
)

var age int

func main() {
	//方式一
	name := flag.String("name", "张三", "please input your name")
	//方式二
	flag.IntVar(&age, "age", 18, "please input your age")

	flag.Parse()
	fmt.Println(*name)
	fmt.Println(age)

	// 命令行输入：./flag -h
}
