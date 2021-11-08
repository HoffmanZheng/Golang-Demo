package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./template_example.tmpl")
	if err != nil {
		fmt.Print("error during parsing template file: ", err)
		return
	}
	user := UserInfo{
		Name:   "李四",
		Gender: "男",
		Age:    18,
	}
	temp.Execute(w, user)
}

func main() {
	http.HandleFunc("/", helloHandleFunc)
	http.ListenAndServe(":80", nil)
}
