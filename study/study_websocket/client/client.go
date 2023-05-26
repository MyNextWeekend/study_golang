package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/"}
	log.Printf("连接到 %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("无法连接到 WebSocket 服务器:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("读取消息错误:", err)
				return
			}
			log.Printf("收到消息: %s", message)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("写入消息错误:", err)
				return
			}
		case <-interrupt:
			log.Println("接收到中断信号，关闭连接...")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("发送关闭消息错误:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(1 * time.Second):
			}
			return
		}
	}
}
