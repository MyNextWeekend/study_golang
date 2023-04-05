package a_creation_mode

import "fmt"

/**
工厂方法模式
*/

// ======= 抽象层 =========

// 水果类(抽象接口)
type Fruit2 interface {
	Show() //接口的某方法
}

// 工厂类(抽象接口)
type AbstractFactory2 interface {
	CreateFruit() Fruit2 //生产水果类(抽象)的生产器方法
}

// ======= 基础类模块 =========
type Apple2 struct {
	Fruit2 //为了易于理解显示继承(此行可以省略)
}

func (apple *Apple2) Show() {
	fmt.Println("我是苹果")
}

type Banana2 struct {
	Fruit2
}

func (banana *Banana2) Show() {
	fmt.Println("我是香蕉")
}

type Pear2 struct {
	Fruit2
}

func (pear *Pear2) Show() {
	fmt.Println("我是梨")
}

// (+) 新增一个"日本苹果"
type JapanApple2 struct {
	Fruit2
}

func (jp *JapanApple2) Show() {
	fmt.Println("我是日本苹果")
}

// ========= 工厂模块  =========
// 具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit2 {
	var fruit Fruit2

	//生产一个具体的苹果
	fruit = new(Apple2)

	return fruit
}

// 具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit2 {
	var fruit Fruit2

	//生产一个具体的香蕉
	fruit = new(Banana2)

	return fruit
}

// 具体的梨工厂
type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit2 {
	var fruit Fruit2

	//生产一个具体的梨
	fruit = new(Pear2)

	return fruit
}

// 具体的日本工厂
type JapanAppleFactory struct {
	AbstractFactory
}

func (fac *JapanAppleFactory) CreateFruit() Fruit2 {
	var fruit Fruit2

	//生产一个具体的日本苹果
	fruit = new(JapanApple2)

	return fruit
}

// ========= 业务逻辑层  =========
func main() {
	/*
			本案例为了突出根据依赖倒转原则与面向接口编程特性。
		    一些变量的定义将使用显示类型声明方式
	*/

	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂
	var appleFac AbstractFactory2
	appleFac = new(AppleFactory)
	//2-生产相对应的具体水果
	var apple Fruit2
	apple = appleFac.CreateFruit()

	apple.Show()

	//需求2：需要一个具体的香蕉对象
	//1-先要一个具体的香蕉工厂
	var bananaFac AbstractFactory2
	bananaFac = new(BananaFactory)
	//2-生产相对应的具体水果
	var banana Fruit2
	banana = bananaFac.CreateFruit()

	banana.Show()

	//需求3：需要一个具体的梨对象
	//1-先要一个具体的梨工厂
	var pearFac AbstractFactory2
	pearFac = new(PearFactory)
	//2-生产相对应的具体水果
	var pear Fruit2
	pear = pearFac.CreateFruit()

	pear.Show()

	//需求4：需要一个日本的苹果？
	//1-先要一个具体的日本评估工厂
	var japanAppleFac AbstractFactory2
	japanAppleFac = new(JapanAppleFactory)
	//2-生产相对应的具体水果
	var japanApple Fruit2
	japanApple = japanAppleFac.CreateFruit()

	japanApple.Show()
}
