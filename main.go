package main

import (
	"log"
	"net/http"
	"ws/routers"
	"ws/service"
)

func main() {
	r := routers.InitRouter()
	go service.Manager.Start()

	err := http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
