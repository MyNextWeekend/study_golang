package study_viper

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
	//添加读取的配置文件路径, （如果从main运行从根目录  如果test方法运行从当前目录）
	vip.AddConfigPath("./config/")
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
