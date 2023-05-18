package study_snyc

import (
	"fmt"
	"sync"
	"testing"
)

//Go1.9版本新增了sync.Map它是原生支持并发安全的map，sync.Map封装了更为复杂的数据结构实现了比之前加读写锁锁map更优秀的性能。
// 查询一个key
//func (m *Map) Load(key interface{}) (value interface{}, ok bool)

// 设置key value
//func (m *Map) Store(key, value interface{})

// 如果key存在则返回key对应的value，否则设置key value
//func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)

// 删除一个key
//func (m *Map) Delete(key interface{})

// 遍历map，仍然是无序的
//func (m *Map) Range(f func(key, value interface{}) bool)

func TestMap(t *testing.T) {
	sMap := &sync.Map{}
	sMap.Store("张三", 19)
	fmt.Println(sMap.Load("张三"))
}
