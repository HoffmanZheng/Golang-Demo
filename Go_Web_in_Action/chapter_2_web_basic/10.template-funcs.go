package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Welcome() string {
	return "Welcome"
}

func Doing(name string) string {
	return name + ", Learning Go Web template "
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./template_func.tmpl")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// define an anonymous template function
	loveGo := func() string {
		return "Welcome learning《Go Web in Action》 together"
	}
	// call Funcs to add template function
	tmpl1, err := template.New("funcs").Funcs(template.FuncMap{"loveGo": loveGo}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	funcMap := template.FuncMap{
		"Welcome": Welcome,
		"Doing":   Doing,
	}
	name := "Hoffman"
	tmpl2, err := template.New("test").Funcs(funcMap).Parse("{{Welcome}}<br/>{{Doing .}}")
	if err != nil {
		panic(err)
	}

	tmpl1.Execute(w, name)
	tmpl2.Execute(w, name)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8087", nil)
}
