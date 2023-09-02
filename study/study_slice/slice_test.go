package study_slice

import (
	"testing"
)

func TestName(t *testing.T) {
	//初始化方法一,长度和容量值一致
	s1 := make([]int, 10)
	t.Logf("s1 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s1, len(s1), cap(s1))
	t.Logf("s1 索引位置1的值为：%v\n\n", s1[1])

	//初始化方法二,长度和容量值不一致,预留一些位置存放数据
	s2 := make([]int, 5, 10)
	t.Logf("s2 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s2, len(s2), cap(s2))
	t.Logf("s2 索引位置1的值为：%v\n\n", s2[1])

	//初始化方法三,手动初始化值
	s3 := []int{1, 2, 3, 4}
	t.Logf("s3 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s3, len(s3), cap(s3))
	t.Logf("s3 索引位置1的值为：%v\n\n", s3[1])

	//初始化方法四,只声明，不初始化（不能直接赋值，可以append）
	var s4 []int
	s4 = append(s4, 1)
	t.Logf("s4 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s4, len(s4), cap(s4))
	t.Logf("s4 索引位置1的值为：%v\n\n", s4[0])

	// 截取一（实际指向同一个底层数组，修改会影响旧切片。扩容会导致数组拷贝，不影响旧切片）
	s5 := []int{1, 2, 3, 4}
	s6 := s5[2:]
	t.Logf("s5 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s5, len(s5), cap(s5))
	t.Logf("s6 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s6, len(s6), cap(s6))
	t.Logf("s6 索引位置1的值为：%v\n\n", s6[0])

	// 截取二（实际指向同一个底层数组，修改会影响旧切片。扩容会导致数组拷贝，不影响旧切片）
	s7 := make([]int, 5, 10)
	s8 := s7[2:]         //此处相当于[2:10]，结尾为容量的最大值，不是长度的最大值
	s8 = append(s8, 666) //s8追加数据的时候没有发生数组拷贝，底层数组被修改了，但是由于s7切片的长度不变，故目前看旧切片不受影响
	t.Logf("s7 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s7, len(s7), cap(s7))
	t.Logf("s8 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s8, len(s8), cap(s8))
	s7 = append(s7, 777) //s7追加数据，修改了s8追加的数据
	t.Logf("s7 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s7, len(s7), cap(s7))
	t.Logf("s8 分配了初始值为：%v ,长度是：%v ,容量是：%v\n", s8, len(s8), cap(s8))
	t.Logf("s8 索引位置1的值为：%v\n\n", s8[0])
}
