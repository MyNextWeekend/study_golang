package main

import "fmt"

func main() {
	var c []int
	c = append(c, 1)
	fmt.Printf("地址是：%p\n", c)
	c = append(c, 1)
	fmt.Printf("地址是：%p\n", c)
	c = append(c, 1)
	fmt.Printf("地址是：%p\n", c)
	c = append(c, 1)
	fmt.Printf("地址是：%p\n", c)
}
