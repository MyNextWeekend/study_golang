package test

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
)

func TestTcp(t *testing.T) {
	listen, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		fmt.Println(conn)
		go func(conn net.Conn) {
			for {
				conn.Write([]byte("报上名来！！"))
				buf := make([]byte, 4096)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("err:", err)
					return
				}
				if err != nil && err != io.EOF {
					fmt.Println("Conn Read err:", err)
					return
				}
				fmt.Println(string(buf[:n-1]))
			}
		}(conn)

	}
}

func TestHttp(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", nil)
}
