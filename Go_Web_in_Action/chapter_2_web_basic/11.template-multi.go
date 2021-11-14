package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}

func tmplSample(w http.ResponseWriter, r *http.Request) {
	// parse multiple template files and fulfill the nested
	tmpl, err := template.ParseFiles("./template_multi.tmpl", "./ul.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := User{
		Name:   "Hoffman",
		Gender: "Male",
		Age:    28,
	}
	tmpl.Execute(w, user)
	fmt.Println(tmpl)
}

func main() {
	http.HandleFunc("/", tmplSample)
	http.ListenAndServe(":8087", nil)
}
