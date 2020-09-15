package main

import (
	"log"
	"net/http"
	"ws/models"
	"ws/pkg/database"
	"ws/routers"
	"ws/service"
)

func main() {
	r := routers.InitRouter()
	go service.Manager.Start()

	database.DB.AutoMigrate(
		&models.ClientMessage{},
	)

	err := http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
