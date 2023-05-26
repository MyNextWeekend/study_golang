package server

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var C *Clients

type Clients struct {
	clientMap sync.Map
}

func init() {
	C = &Clients{}
}

func (c *Clients) Add(key string) error {
	_, ok := c.clientMap.Load(key)
	if ok {
		return errors.New("信息已存在，请勿重复提交！！！")
	}

	client := NewClient(10 * time.Second)
	err := client.Conn()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 加入到map中
	c.clientMap.Store(key, client)
	return nil
}

func (c *Clients) Stop(key string) error {
	value, ok := c.clientMap.Load(key)
	if !ok {
		return errors.New("输入的信息不存在")
	}
	// 调用stop方法
	value.(*Client).Stop()
	// 从map中删除key
	c.clientMap.Delete(key)
	return nil
}

func (c *Clients) QueryAll() (result []string) {
	result = make([]string, 0)
	c.clientMap.Range(func(key, value any) bool {
		result = append(result, key.(string))
		return true
	})
	return result
}
