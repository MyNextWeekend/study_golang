package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", handleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("升级为 WebSocket 连接失败:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息错误:", err)
			break
		}

		log.Printf("收到消息: %s", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("发送消息错误:", err)
			break
		}
	}
}
