package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"ws/service"
)

// 连接websocket
func WebsocketServe(c *gin.Context) {
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 连接时的message
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	// 鉴权
	auth := &service.Auth{}
	err = json.Unmarshal(message, auth)
	if err != nil {
		log.Println(err)
		return
	}

	err = auth.Authorized()
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}

	client := &service.Client{
		ClientID:    auth.Uid,
		Conn:        conn,
		Message:     make(chan []byte, 256),
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
		GroupList:   auth.Group,
	}

	service.Manager.ConnectChan <- client
}
