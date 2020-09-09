package routers

import (
	"github.com/gin-gonic/gin"
	"ws/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/send", api.Send)
	r.GET("/ws", api.WebsocketServe)

	return r
}
