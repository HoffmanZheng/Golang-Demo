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
	uid, _ := strconv.Atoi(query["uid"][0])

	//此处调用模型从数据库中获取数据
	user := model.GetUser(uid)
	fmt.Println(user)

	t, _ := template.ParseFiles("view/user.html")
	userInfo := []string{user.Name, user.Phone}
	t.Execute(w, userInfo)
}
