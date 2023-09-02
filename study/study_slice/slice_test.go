package study_slice

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//初始化方法一,长度和容量值一致
	s1 := make([]int, 10)
	fmt.Printf("s1 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s1, len(s1), cap(s1))
	fmt.Printf("s1 索引位置1的值为：%v\n", s1[1])

	//初始化方法二,长度和容量值不一致,预留一些位置存放数据
	s2 := make([]int, 5, 10)
	fmt.Printf("s2 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s2, len(s2), cap(s2))
	fmt.Printf("s2 索引位置1的值为：%v\n", s2[1])

	//初始化方法三,手动初始化值
	s3 := []int{1, 2, 3, 4}
	fmt.Printf("s3 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s3, len(s3), cap(s3))
	fmt.Printf("s3 索引位置1的值为：%v\n", s3[1])

	//初始化方法四,只声明，不初始化（不能直接赋值，可以append）
	var s4 []int
	s4 = append(s4, 1)
	fmt.Printf("s4 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s4, len(s4), cap(s4))
	fmt.Printf("s4 索引位置1的值为：%v\n", s4[0])

}
