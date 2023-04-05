package a_creation_mode

import "fmt"

/**
简单工厂模式
*/

// ======= 抽象层 =========

// 水果类(抽象接口)
type Fruit1 interface {
	Show() //接口的某方法
}

// ======= 基础类模块 =========

type Apple1 struct {
	Fruit1 //为了易于理解显示继承(此行可以省略)
}

func (apple *Apple1) Show() {
	fmt.Println("我是苹果")
}

type Banana1 struct {
	Fruit1
}

func (banana *Banana1) Show() {
	fmt.Println("我是香蕉")
}

type Pear1 struct {
	Fruit1
}

func (pear *Pear1) Show() {
	fmt.Println("我是梨")
}

// ========= 工厂模块  =========
// 一个工厂， 有一个生产水果的机器，返回一个抽象水果的指针
type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruit1 {
	var fruit Fruit1

	if kind == "apple" {
		fruit = new(Apple1)
	} else if kind == "banana" {
		fruit = new(Banana1)
	} else if kind == "pear" {
		fruit = new(Pear1)
	}

	return fruit
}

// ==========业务逻辑层==============
func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
