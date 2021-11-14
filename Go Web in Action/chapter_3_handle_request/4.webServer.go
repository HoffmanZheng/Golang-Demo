package main

import (
	"log"
	"net/http"

	controller "github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_3_handle_request/controller"
)

func main() {
	http.HandleFunc("/getUser", controller.UserController{}.GetUser)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}
