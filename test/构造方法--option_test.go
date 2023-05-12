package test

import (
	"fmt"
	"testing"
	"time"
)

/**
option 模式：
	可以灵活配置可选参数与必填参数
*/

type Option func(server *Server)

func WithProtocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConns(maxConns int) Option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

func NewUser(addr string, port int, options ...Option) *Server {
	// 先初始化一个带默认值的对象
	ser := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  time.Second,
		MaxConns: 10,
	}

	// 通过同一个保重暴漏的闭包函数来给对象特定值赋值处理
	for _, option := range options {
		option(&ser)
	}

	return &ser
}

func TestNewUser(t *testing.T) {
	// 使用的时候传入对应方法 即可覆盖默认值
	server1 := NewUser("192.168.0.1", 80, WithTimeout(time.Second*3), WithMaxConns(30))
	server2 := NewUser("192.168.0.2", 90, WithTimeout(time.Second*8), WithProtocol("udp"))
	fmt.Println(server1)
	fmt.Println(server2)
}
