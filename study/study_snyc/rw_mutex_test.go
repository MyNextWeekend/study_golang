package study_snyc

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type Bank2 struct {
	sync.RWMutex // 结构体内部维护一个锁
	balance      map[string]float64
}

//收入
func (b *Bank2) In2(account string, value float64) {
	//加锁，保证同一时间只有一个协程能访问下面的代码
	b.Lock()
	defer b.Unlock()

	if _, ok := b.balance[account]; !ok {
		b.balance[account] = 0.0
	}
	b.balance[account] += value

}

func (b *Bank2) Out2(account string, value float64) error {
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

func (b *Bank2) Query(account string) float64 {
	//查询使用读锁，没有写入的情况下，不影响读
	b.RLock()
	defer b.RUnlock()

	if _, ok := b.balance[account]; !ok {
		return 0.0
	}
	return b.balance[account]
}

func TestBank2(t *testing.T) {
	bank2 := Bank2{balance: map[string]float64{}}
	bank2.In2("z", 100)
	bank2.In2("z", 100)
	fmt.Println(bank2.Query("z"))
}
