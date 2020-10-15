package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ws/models"
	"ws/pkg/app"
	"ws/pkg/database"
)

type addClientFrom struct {
	ClientID string `json:"client_id" valid:"Required"`
	GroupID  uint   `json:"group_id" valid:"Required"`
}

type CreateGroupForm struct {
	Name string `json:"name" valid:"Required"`
}

func CreateGroup(c *gin.Context) {
	form := new(CreateGroupForm)

	if err := app.BindAndValid(c, form); err != nil {
		c.JSON(http.StatusPaymentRequired, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&models.Group{
		Name: form.Name,
	})

	c.Status(http.StatusCreated)
}

func DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete()
}

func AddClient2Group(c *gin.Context) {

}

func RemoveClientFromGroup(c *gin.Context) {

}
