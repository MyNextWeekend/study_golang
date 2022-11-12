package main

import "fmt"

func main() {
	c := make([][]int, 0)
	a := []int{1, 2, 3}
	c = append(c, a)
	fmt.Println(c)
	a = append(a, 888)
	c = append(c, a)
	fmt.Println(c)
}
