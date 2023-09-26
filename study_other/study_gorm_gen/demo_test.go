package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"study_golang/study_other/study_gorm_gen/dal"
	"study_golang/study_other/study_gorm_gen/model"
	"testing"
	"time"
)

func InitDb() {
	dsn := "study_golang:6xkiKTGzfPE7bXWE@(106.55.186.222:3306)/study_golang?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("open db err: " + err.Error())
	}
	dal.SetDefault(db) // 设置数据库链接
}

func TestCreat(t *testing.T) {
	InitDb()

	user := model.UserInfo{
		NickName: "张三",
		UserName: "zhangsan",
		Password: "12345",
		Role:     "admin",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err := dal.UserInfo.Create(&user)
	if err != nil {
		fmt.Println("creat err: " + err.Error())
	}
}

func TestQuery(t *testing.T) {
	InitDb()

	users, err := dal.UserInfo.Where(dal.UserInfo.UserName.Like("z%")).Find()
	if err != nil {
		fmt.Println("creat err: " + err.Error())
	}
	for i, user := range users {
		fmt.Printf("序号：%v 的值是:%v\n", i, user)
	}
}
