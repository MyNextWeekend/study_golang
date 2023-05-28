package server

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/url"
	"strings"
	"study_golang/example/config"
	"time"
)

type Client struct {
	agentId string
	end     chan struct{}   // 结束标识
	conn    *websocket.Conn //websocket链接
	ticker  *time.Ticker    // 间隔固定时间向服务器发送ping，保持心跳
	s       *server
}

func NewClient(agentId string, s *server) *Client {
	return &Client{
		agentId: agentId,
		end:     make(chan struct{}),
		ticker:  time.NewTicker(10 * time.Second),
		s:       s,
	}
}

func (c *Client) Conn() error {
	path := config.WebSocketPath
	u := url.URL{Scheme: config.WebSocketScheme, Host: config.WebSocketHost, Path: path}
	fmt.Println("url：" + u.String())

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return errors.New("conn websocket failed " + err.Error())
	}
	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)
	// 判断是否建立成功，如果已经建立链接
	if strings.Contains(string(result), "别处") {
		return errors.New("已经建立连接")
	}
	c.conn = conn

	go c.readMessage()  // 启动一个读协程
	go c.writeMessage() // 启动一个写的协程
	return nil
}

// 读返回的消息
func (c *Client) readMessage() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息错误:", err)
			return
		}
		fmt.Printf("收到消息:%s\n", message)
	}
}

// 定时写消息
func (c *Client) writeMessage() {
	for {
		select {
		case <-c.end:
			fmt.Println("接收到中断信号，关闭连接...")
			err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("发送关闭消息错误:", err)
			}
			return
		case <-c.ticker.C:
			err := c.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				fmt.Println("写入消息错误:", err)
				c.Stop()
			}
		}
	}
}

func (c *Client) Stop() {
	fmt.Println("stop ", c.agentId)
	close(c.end)
	// 从map中删除key
	c.s.clientMap.Delete(c.agentId)
}
