package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	DbUser        string
	DbPassWord    string
	DbName        string
	RedisDb       string
	RedisAddr     string
	RedisPassWord string
	RedisDbName   string
)

func main() {
	file, err := ini.Load("./study_other/study_ini/config.ini")
	if err != nil {
		fmt.Println("ini load failed", err.Error())
		return
	}
	LoadMysql(file)
	LoadRedis(file)
}

func LoadMysql(file *ini.File) {
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPassWord = file.Section("redis").Key("RedisPassWord").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
