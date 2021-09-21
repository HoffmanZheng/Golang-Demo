package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	engine.POST("/user/login", goLogin)
	engine.Run()
}

func goLogin(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	c.String(200, "username=%s, password+%s", name, password)
}
