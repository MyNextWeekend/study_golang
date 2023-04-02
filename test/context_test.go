package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("收到结束信息\n")
			return
		default:
			fmt.Print("干活1\n")
			time.Sleep(3 * time.Second)
			//fmt.Print("干活2\n")
			//time.Sleep(3 * time.Second)
			//fmt.Print("干活3\n")
			//time.Sleep(3 * time.Second)
			//fmt.Print("干活4\n")
			//time.Sleep(3 * time.Second)
			//fmt.Print("干活5\n")
			//time.Sleep(3 * time.Second)
		}
	}
}
func TestContext(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	//defer cancelFunc()
	go work(ctx)
	time.Sleep(6 * time.Second)
	cancelFunc()
	fmt.Print("主动结束操作")
	time.Sleep(10 * time.Second)
	fmt.Print("done")
}
