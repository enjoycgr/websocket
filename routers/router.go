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

	return r
}
