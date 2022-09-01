package main

import "fmt"

type animal interface {
	speak()
	move()
}
type dog struct {
}

func (d *dog) speak() {
	fmt.Println("汪汪汪")
}

func (d *dog) move() {
	fmt.Println("动起来")
}
func (d *dog) sing() {
	fmt.Println("唱歌")
}

func main() {
	var do dog
	do.speak()
	do.move()
	do.sing()
	var abc animal
	abc = &do
	abc.speak()
	abc.move()

}
