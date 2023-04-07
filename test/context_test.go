package test

import (
	"context"
	"fmt"
	"time"
)

/**
context 的使用规则:
	context 应该作为函数参数传递，而不是 struct 的一个 field
	context 应该是函数的第一个参数
	不要传递 nil context，即使是不使用，如果不知道用啥，用 context.TODO 吧
	只使用 context 传递请求上下文，而不是为了传递可选参数
	context 可能会被同一个函数在不同的 goroutine 中使用，他是并发安全的
*/

// 程序在子协程运行了30s后就会被主协程关闭。
func CancelTest() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("协程退出，停止了...")
				return
			default:
				fmt.Println("协程运行中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 30)
	fmt.Println("两分钟时间到了，关闭子协程")
	cancel()
	time.Sleep(time.Second * 10)
	fmt.Println("演示结束")
}

// 子协程会在运行1分钟后主动退出，然后接着10秒后主协程也会运行完退出
func TimeOutTest() {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("协程退出，停止了...")
				return
			default:
				fmt.Println("协程运行中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 70)
	fmt.Println("演示结束")
}

// context.WithDeadline和context.WithTimeout这两个方法的功能是相似的，
// 区别是WithDeadline可以在指定任意时刻调用cancel函数，而WithTimeout只能基于调用时的时间然后加上一段时间调用cancel函数。
// 函数中context.WithDeadline需要传入的是个时间点，我写的时间点时time.Now().Add(time.Minute * 1)，
// 当前时间后的一分钟触发cancel函数(你可以选择任意时间点)，所以子协程会在运行一分钟后结束运行，
// 当然未到一分钟时，你可以主动调用withDeadline方法所返回的cancel函数对子协程提前管理。
func DeadLineTest() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*1))
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("协程退出，停止了...")
				return
			default:
				fmt.Println("协程运行中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 70)
	fmt.Println("演示结束")
}

// 将配置好的key-value的context传入子协程，然后子协程可以使用context.value()方法访问到值，这在各个父子关资的协程中传值比较方便。
func WithValueTest() {
	ctx, cancel := context.WithCancel(context.Background()) //附加值
	valueCtx := context.WithValue(ctx, "test", "子协程1")
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): //取出值
				fmt.Println(ctx.Value("test"), "监控退出，停止了...")
				return
			default: //取出值
				fmt.Println(ctx.Value("test"), "goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel() //为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
