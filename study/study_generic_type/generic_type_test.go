package study_generic_type

import (
	"fmt"
	"testing"
)

func rangeSlice[t any](slice []t) {
	for i, t2 := range slice {
		fmt.Printf("索引位置：%v,值为：%v。 ", i, t2)
	}
	fmt.Println("done")
}

func TestName(t *testing.T) {
	s1 := []int{1, 2, 3}
	rangeSlice(s1)

	s2 := []string{"a", "b", "c"}
	rangeSlice[string](s2) //调用方法的时候可以手动指定范型类型，也可省略

}
