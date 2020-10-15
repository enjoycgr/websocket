package routers

import (
	"github.com/gin-gonic/gin"
	"ws/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/send2client", api.Send2Client)
	r.POST("/send2group", api.Send2Group)
	r.GET("/ws", api.WebsocketServe)

	r.POST("/group", api.CreateGroup)
	r.DELETE("/group/:id", api.DeleteGroup)
	r.POST("/group/:id/client", api.AddClient2Group)
	r.DELETE("/group/:id/client", api.RemoveClientFromGroup)

	r.GET("/message/client", api.ClientMessage)
	//r.GET("/message/group", api.GroupMessage)
	return r
}
