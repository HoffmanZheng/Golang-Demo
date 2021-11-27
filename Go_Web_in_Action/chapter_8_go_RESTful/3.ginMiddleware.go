package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		// logic call before handlerfunc
		ctx.Set("example", "hi, there is a middleware data")

		ctx.Next() // call next middleware or handlerFunc

		latency := time.Since(t)
		log.Print("hanle time of request: ", latency)

		status := ctx.Writer.Status()
		log.Print("response status: ", status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())       // log middleware to log request
	r.Use(gin.Recovery()) // intercept panic to aviod program crash
	r.GET("/hi", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)
		log.Println(example)
	})
	r.Run(":8080")
}
