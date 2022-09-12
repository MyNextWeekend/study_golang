package main

type student struct {
	name string
	age  chan int
}

func main() {

	//未显示初始化的对象，自动初始化了
	//var abc student
	//fmt.Println(&abc == nil)    //false
	//fmt.Println(abc.age == nil) //true
	//fmt.Println(abc)

	//var m map[string]string
	//fmt.Println(m == nil) //true

	//m := new(map[string]int)
	//m[""] = 1
	//fmt.Println(m)

	//未初始化的通道，不能写入
	//var c chan int
	//c <- 1 // 死锁
	//fmt.Println("end")

	// 未初始化可以添加元素
	//var sli []int
	//fmt.Println(sli == nil) //true
	//sli = append(sli, 1)
	//fmt.Println(sli)

	// 未初始化不能使用
	//var m map[string]string
	//m["no1"] = "001"
	//fmt.Println(m)

}
