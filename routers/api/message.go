package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ws/models"
	"ws/pkg/app"
	"ws/pkg/database"
	"ws/service"
)

type sendData struct {
	Message string `json:"message" valid:"Required"`
	Sender  string `json:"sender" valid:"Required"`
}

type groupData struct {
	sendData
	Groups []uint `json:"groups" valid:"Required"`
}

type clientData struct {
	sendData
	Clients []string `json:"clients" valid:"Required"`
}

// 发送消息给group
func Send2Group(c *gin.Context) {
	input := new(groupData)

	if err := app.BindAndValid(c, input); err != nil {
		c.JSON(http.StatusPaymentRequired, gin.H{"message": err.Error()})
		return
	}

	var groupMessageRead []models.GroupMessageRead
	sent := make(map[string]bool)

	for _, v := range input.Groups {
		// 保存消息
		groupMessage := models.GroupMessage{
			GroupID: v,
			Message: input.Message,
			Sender:  input.Sender,
		}
		database.DB.Create(&groupMessage)

		for _, client := range service.Manager.Group[v] {
			// 判断重复发送
			if _, ok := sent[client.ClientID]; ok == false {
				client.Message <- []byte(input.Message)
				sent[client.ClientID] = true

				groupMessageRead = append(groupMessageRead, models.GroupMessageRead{
					GroupMessageID: groupMessage.ID,
					ClientID:       client.ClientID,
				})
			}
		}

		// 保存已读消息
		database.DB.Create(groupMessageRead)
	}
}

// 发送消息给client
func Send2Client(c *gin.Context) {
	input := new(clientData)

	if err := app.BindAndValid(c, input); err != nil {
		c.JSON(http.StatusPaymentRequired, gin.H{"message": err.Error()})
		return
	}

	var clientMessage []models.ClientMessage

	for _, v := range input.Clients {
		m := models.ClientMessage{
			ClientID: v,
			Message:  input.Message,
			Sender:   input.Sender,
		}

		if client, ok := service.Manager.ClientMap[v]; ok {
			client.Message <- []byte(input.Message)
			m.Read = 1
		}

		clientMessage = append(clientMessage, m)
	}

	// 保存消息
	database.DB.Create(clientMessage)
}

// 客户端消息列表
func ClientMessage(c *gin.Context) {

}
