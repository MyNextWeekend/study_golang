package study_snyc

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
sync.Cond指的是同步条件变量，一般需要与互斥锁组合使用，本质上是一些正在等待某个条件的协程的同步机制。

sync.Cond有3个函数Wait、Signal、Broadcast
func (c *Cond) Wait()		// Wait 等待通知
func (c *Cond) Signal()		// Signal 单播通知
func (c *Cond) Broadcast()	// Broadcast 广播通知
*/

var sharedRsc = make(map[string]interface{})

func TestCond(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	cond := sync.NewCond(&sync.Mutex{})

	go func() {
		defer wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()

		for len(sharedRsc) == 0 { // 不能假设条件为真。
			fmt.Println("goroutine 01 is wait...")
			// Wait()会自动释放c.L，并挂起调用者的goroutine。之后恢复执行，Wait()会在返回时对c.L加锁。
			// 除非被Signal或者Broadcast唤醒，否则Wait()不会返回。
			cond.Wait()
			fmt.Println("goroutine 01 is go go go...")
		}
		fmt.Println("goroutine 01 working", sharedRsc["rsc1"])
	}()

	go func() {
		defer wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()

		for len(sharedRsc) == 0 { // 不能假设条件为真。
			cond.Wait()
		}
		fmt.Println("goroutine 02 working", sharedRsc["rsc2"])
	}()

	time.Sleep(3 * time.Second)
	cond.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	cond.Broadcast()
	cond.L.Unlock()
	wg.Wait()
}
