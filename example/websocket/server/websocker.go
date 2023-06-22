package server

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	config2 "study_golang/example/websocket/config"
	"time"
)

type Client struct {
	agentId string
	end     chan struct{}   // 结束标识
	conn    *websocket.Conn //websocket链接
	ticker  *time.Ticker    // 间隔固定时间向服务器发送ping，保持心跳
	s       *server
}

func getUrlPath(agentId string) (urlPath string) {
	org := "10000030"
	phoneNum := "6543"
	key := fmt.Sprintf("7A89843D507A3229745CFB23C2B27871%s%s%s", agentId, org, phoneNum)
	ag := fmt.Sprintf("%x", sha256.Sum256([]byte(key)))
	urlPath = fmt.Sprintf("outMessageSocket/%s/%s/%s/%s", ag, agentId, org, phoneNum)
	return
}

func NewClient(agentId string, s *server) *Client {
	return &Client{
		agentId: agentId,
		end:     make(chan struct{}),
		ticker:  time.NewTicker(10 * time.Second),
		s:       s,
	}
}

func (c *Client) Conn() (err error) {
	path := config2.WebSocketPath + getUrlPath(c.agentId)
	u := url.URL{Scheme: config2.WebSocketScheme, Host: config2.WebSocketHost, Path: path}
	config2.Logger.Infof("用户：%s --> url：%s\n", c.agentId, u.String())

	c.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return errors.New("conn websocket failed " + err.Error())
	}

	go c.readMessage()  // 启动一个读协程
	go c.writeMessage() // 启动一个写的协程
	return nil
}

// 读返回的消息
func (c *Client) readMessage() {
	for {
		select {
		case <-c.end:
			config2.Logger.Infof("%s close read line success\n", c.agentId)
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				config2.Logger.Warn("read message err ", err.Error())
				c.Stop() // 读异常就关闭链接
				return
			}
			config2.Logger.Infof("%s get massage: %s\n", c.agentId, message)
		}
	}
}

// 定时写消息
func (c *Client) writeMessage() {
	for {
		select {
		case <-c.end:
			_ = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			config2.Logger.Infof("%s close write line success\n", c.agentId)
			return
		case <-c.ticker.C:
			err := c.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				config2.Logger.Warnf("%s write message err: %s\n", c.agentId, err.Error())
				c.Stop() // 写入异常就关闭链接
				return
			}
		}
	}
}

func (c *Client) Stop() {
	select {
	case <-c.end:
		return
	default:
		close(c.end)
		// 从map中删除key
		c.s.clientMap.Delete(c.agentId)
	}
}
