package study_redis

import (
	"context"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	InitRedis()

	ctx := context.Background()

	err := Set(ctx, "zhangsan", "lisi")
	if err != nil {
		fmt.Println("redis set err: " + err.Error())
	}

	result, err := Get(ctx, "zhangsan")
	if err != nil {
		fmt.Println("redis get key err: " + err.Error())
		return
	}

	fmt.Println("从redis中获取的结果是：" + result)
	fmt.Println("done")
}
