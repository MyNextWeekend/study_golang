package study_viper

import (
	"fmt"
	"testing"
)

func TestVip(t *testing.T) {
	InitViper()

	//fmt.Println("name:", vip.GetString("name"))
	//fmt.Println("age", vip.GetInt("age"))

	fmt.Println("ip", vip.GetString("mysql.ip"))
	fmt.Println("port", vip.GetInt("mysql.port"))
	fmt.Println("user", vip.GetString("mysql.user"))
	fmt.Println("password", vip.GetString("mysql.password"))
}
