package main

import (
	"fmt"
	"study_golang/study_other/study_gorm/dao"
	"time"
)

// 插入一条数据
func InsertOne() {
	user := dao.User{
		Name:     "张三",
		Email:    "123@123.com",
		Age:      18,
		Birthday: time.Now(),
	}
	result := DB.Create(&user)
	/**
	user.ID             // 返回插入数据的主键
	result.Error        // 返回 error
	result.RowsAffected // 返回插入记录的条数
	*/
	if result.Error != nil {
		fmt.Println("Insert Databases Error:" + result.Error.Error())
		return
	}
	fmt.Println("插入的数据id是：", user.ID)
	fmt.Println("数据：", user)
	fmt.Println("成功的条数：", result.RowsAffected)
}

// 插入多条数据
func InsertMany() {
	users := []dao.User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	result := DB.Create(&users)
	if result.Error != nil {
		fmt.Println("Insert Databases Error:" + result.Error.Error())
		return
	}
	fmt.Println("成功的条数：", result.RowsAffected)
	for _, user := range users {
		fmt.Println("插入的数据id是：", user.ID)
		//fmt.Println("数据：", user)
	}
}

// 查询一条不带条件
func SelectOne() {
	var user dao.User
	// 获取第一条记录（主键升序）
	//result := DB.First(&user)
	// 获取最后一条记录（主键降序）
	result := DB.Last(&user)
	if result.Error != nil {
		fmt.Println("Select Databases Error:" + result.Error.Error())
		return
	}
	fmt.Println("查询出来的数据：", user)
}

// 查询多条带条件
func SelectMany() {
	var user []dao.User
	// 获取所有数据
	result := DB.Where("age = ?", 18).Find(&user)
	if result.Error != nil {
		fmt.Println("Select Databases Error:" + result.Error.Error())
		return
	}
	fmt.Println("查询出来的条数：", result.RowsAffected)
	fmt.Println("查询出来的数据：", user)
}

func main() {
	InitORM()
	//CreatTable()
	//InsertOne()
	//InsertMany()
	//SelectOne()
	SelectMany()
}
