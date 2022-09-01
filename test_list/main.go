package main

import "fmt"

func main() {
	numList := []int{1, 2, 3}
	//遍历切片
	for i, i2 := range numList {
		fmt.Println(i, i2)
	}
	//切片增加
	numList = append(numList, 9)
	//修改元素
	numList[2] = 10
	//删除元素
	//numList = numList[1:3]
	numList = append(numList[0:1], numList[2])
	fmt.Println(numList)
}
