package server

import (
	"errors"
	"fmt"
	"sync"
)

var Server *server

type server struct {
	clientMap sync.Map
}

func init() {
	Server = &server{}
}

func (s *server) AddClient(agentId string) error {
	_, ok := s.clientMap.Load(agentId)
	if ok {
		return errors.New("信息已存在，请勿重复提交！！！")
	}
	client := NewClient(agentId, s)
	err := client.Conn()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 加入到map中
	s.clientMap.Store(agentId, client)
	return nil
}

func (s *server) StopClient(agentId string) error {
	value, ok := s.clientMap.Load(agentId)
	if !ok {
		return errors.New("输入的信息不存在")
	}
	// 调用stop方法
	value.(*Client).Stop()
	return nil
}

func (s *server) QueryAllClient() (result []string) {
	result = make([]string, 0)
	s.clientMap.Range(func(key, value any) bool {
		result = append(result, key.(string))
		return true
	})
	return result
}
