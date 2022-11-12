package main

import "fmt"

type Proper struct {
	Name string
	Data map[interface{}]interface{}
}

func NewProper() *Proper {
	return &Proper{
		Name: "张三",
		Data: make(map[interface{}]interface{}),
	}
}

func main() {

	a := NewProper()
	b := NewProper()
	fmt.Println(a == b)
	fmt.Println(a, b)
	a.Data['a'] = 'a'
	a.Name = "李四"
	fmt.Println(b.Data)
	fmt.Println(b.Name)

}
