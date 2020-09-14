package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"ws/pkg/app"
	"ws/service"
)

type sendGroupData struct {
	Message string   `json:"message" valid:"Required"`
	Group   []string `json:"group" valid:"Required"`
}

type sendClientData struct {
	Message string   `json:"message" valid:"Required"`
	Client  []string `json:"client" valid:"Required"`
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

func Send2Group(c *gin.Context) {
	input := new(sendGroupData)

	if err := app.BindAndValid(c, input); err != nil {
		c.JSON(http.StatusPaymentRequired, gin.H{"message": err.Error()})
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

func Send2Client(c *gin.Context) {
	input := new(sendClientData)

	if err := app.BindAndValid(c, input); err != nil {
		c.JSON(http.StatusPaymentRequired, gin.H{"message": err.Error()})
		return
	}

	for _, v := range input.Client {
		if client, ok := service.Manager.ClientMap[v]; ok {
			client.Message <- []byte(input.Message)
		}
	}
}
