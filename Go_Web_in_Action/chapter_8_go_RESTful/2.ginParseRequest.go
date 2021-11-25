package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Phone string `json:"phone" form:"phone"`
	Age   string `json:"age" form:"age"`
}

// get string from post form

func parsePostForm(ctx *gin.Context) {
	name := ctx.PostForm("name")
	fmt.Printf("get name string from post form: %s\n", name)
}

func parseDefaultPostForm(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "hoffman")
	fmt.Printf("get name string from post form: %s\n", name)
}

func getIdFromPath(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("get id from path: %s \n", id)
}

func parseUserStruct(ctx *gin.Context) {
	u := User{}
	if ctx.ShouldBind(&u) == nil {
		log.Println("phone: " + u.Phone)
		log.Println("age: " + u.Age)
	}
	ctx.String(200, "success")
}

func returnString(ctx *gin.Context) {
	ctx.String(200, "hello, %s", "hoffman")
}

func returnJson(ctx *gin.Context) {
	u := &User{
		Phone: "166666666",
		Age:   "12",
	}
	ctx.JSON(200, u)
}

func main() {
	r := gin.Default()
	r.POST("/getName", parsePostForm)
	r.POST("/getName2", parseDefaultPostForm)
	r.GET("/user/:id", getIdFromPath)
	r.POST("/user/:id", parseUserStruct)
	r.GET("/returnString", returnString)
	r.GET("/returnJson", returnJson)
	r.Run()
}
