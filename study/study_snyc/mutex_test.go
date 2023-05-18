package study_snyc

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type Bank struct {
	sync.Mutex // 结构体内部维护一个锁
	balance    map[string]float64
}

//收入
func (b *Bank) In(account string, value float64) {
	//加锁，保证同一时间只有一个协程能访问下面的代码
	b.Lock()
	defer b.Unlock()

	if _, ok := b.balance[account]; !ok {
		b.balance[account] = 0.0
	}
	b.balance[account] += value

}

func (b *Bank) Out(account string, value float64) error {
	//加锁，保证同一时间只有一个协程能访问下面的代码
	b.Lock()
	defer b.Unlock()

	v, ok := b.balance[account]
	if !ok || v < value {
		return errors.New(account + " not enough balance")
	}
	b.balance[account] -= value
	return nil
}

func TestBank(t *testing.T) {
	bank := Bank{balance: map[string]float64{}}
	bank.In("张三", 13)
	bank.In("李四", 13)
	bank.In("张三", 13)
	err := bank.Out("张三", 26)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(bank.balance)
}
