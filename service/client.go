package service

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 5 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	ClientId    string          // 标识ID
	Conn        *websocket.Conn // 用户连接
	Message     chan []byte
	ConnectTime uint64 // 首次连接时间
	IsDeleted   bool   // 是否删除或下线
	Extend      string // 扩展字段，用户可以自定义
	GroupList   []string
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		Manager.DisconnectChan <- c
	}()

	for {
		select {
		case message, ok := <-c.Message:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func (c *Client) ReadPump() {
	go func() {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		for {
			messageType, message, err := c.Conn.ReadMessage()
			if err != nil {
				if messageType == -1 && websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
					Manager.DisconnectChan <- c
					return
				}
			}

			if string(message) == "ping" {
				c.Message <- []byte("pong")
			}
		}
	}()
}
