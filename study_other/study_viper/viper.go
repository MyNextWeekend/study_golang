package main

import (
	"fmt"
	"github.com/spf13/viper"
)

var vip *viper.Viper

func InitViper() {
	//读取yaml文件
	vip = viper.New()
	//设置读取的配置文件
	vip.SetConfigName("config2")
	//添加读取的配置文件路径
	vip.AddConfigPath("./study_other/study_viper/config/")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	//vip.AddConfigPath("$GOPATH/src/")
	//设置配置文件类型
	vip.SetConfigType("toml")

	if err := vip.ReadInConfig(); err != nil {
		fmt.Println("viper init err", err.Error())
		panic("viper init err" + err.Error())
	}
	fmt.Println("success init viper")

}
