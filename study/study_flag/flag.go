package main

import (
	"flag"
	"fmt"
	"os"
)

var age int

func main() {
	//方式一
	name := flag.String("name", "", "please input your name")
	//方式二
	flag.IntVar(&age, "age", 18, "please input your age")

	flag.Parse()
	// 如果名字为空，就报错
	if *name == "" {
		fmt.Println("Please enter the correct value ！")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(*name)
	fmt.Println(age)

	// 命令行输入：./flag -h
}
