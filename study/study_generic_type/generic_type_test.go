package study_generic_type

import (
	"fmt"
	"testing"
)

// 泛型函数
func rangeSlice[t any](slice []t) {
	for i, t2 := range slice {
		fmt.Printf("索引位置：%v,值为：%v。 ", i, t2)
	}
	fmt.Println("done")
}

func Test01(t *testing.T) {
	s1 := []int{1, 2, 3}
	rangeSlice(s1)

	s2 := []string{"a", "b", "c"}
	rangeSlice[string](s2) //调用方法的时候可以手动指定范型类型，也可省略

}

// 泛型结构体
type Proper[T any] struct {
	name string
	age  int
	data T
}

// 泛型方法
func (p *Proper[T]) getName() string {
	return p.name
}

func Test02(t *testing.T) {
	p := &Proper[string]{
		name: "张三",
		age:  18,
		data: "hello",
	}
	fmt.Println(p.getName())

}
