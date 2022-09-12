package main

import "fmt"

type student struct {
	name string
}

func swapInt(a *int, b *int) {
	a = b
}
func swap(s1 *student, s2 *student) {
	*s1 = *s2
	fmt.Println(s1)
}
func main() {
	s1 := student{"张三"}
	s2 := student{"李四"}
	fmt.Println("start:", &s1)
	fmt.Println("start--:", s1)
	//调用方法，方法会开辟一个内存，操作指针对应的值才会影响外部的值
	swap(&s1, &s2)
	fmt.Println("end:", &s1)
	fmt.Println("end--:", s1)
}
