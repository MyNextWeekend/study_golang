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
	a.Data['a'] = 'a'
	fmt.Print(b.Name)

}
