package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	model "github.com/HoffmanZheng/Golang-Demo/Go_Web_in_Action/chapter_3_handle_request/model"
)

type UserController struct {
}

func (c UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	uid, _ := strconv.Atoi(query["uid"][0]) //  /getUser?uid=1

	// extract data from database
	user := model.GetUser(uid)
	fmt.Println(user)

	t, err := template.ParseFiles("view/user.tmpl")
	if err != nil {
		fmt.Printf("error during template.ParseFiles: %v \n", err)
	}
	t.Execute(w, user)
}
