package a_creation_mode

import (
	"fmt"
	"sync"
)

/**
单例模式
*/

var instance *Instance

type Instance struct {
	name string
}

func NewInstance() {
	instance = &Instance{"go"}
}

func (i *Instance) print() {
	fmt.Println(i.name)
}

func main() {
	var once sync.Once
	once.Do(NewInstance)
	instance.print()
}
