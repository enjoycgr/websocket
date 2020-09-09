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

type InputData struct {
	Group   []string `json:"group" binding:"required"`
	Message string   `json:"message" binding:"required"`
}

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
		ClientId:    auth.Uid,
		Conn:        conn,
		Message:     make(chan []byte, 256),
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
		GroupList:   auth.Group,
	}

	service.Manager.ConnectChan <- client
}

func Send(c *gin.Context) {
	input := new(InputData)
	if err := c.ShouldBind(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sent := make(map[string]bool)
	for _, v := range input.Group {
		for _, client := range service.Manager.Group[v] {
			if _, ok := sent[client.ClientId]; ok == false {
				client.Message <- []byte(input.Message)
				sent[client.ClientId] = true
			}
		}
	}
}
