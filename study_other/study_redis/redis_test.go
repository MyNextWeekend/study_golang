package study_redis

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	InitRedis()
	err := Setex("zhangsan", "lisi", 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	result, err := Get("zhangsa1n")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
