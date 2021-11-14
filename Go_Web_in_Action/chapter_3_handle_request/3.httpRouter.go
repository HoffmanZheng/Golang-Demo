package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HostMap map[string]http.Handler

func (hs HostMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//根据域名获取对应的Handler路由，然后调用处理（分发机制）
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("error")
}

func main() {
	router := httprouter.New()

	// deel with error using PanicHandler
	router.GET("/panic", Index)
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error:%s", v)
	}

	// RESTFul match
	router.GET("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default get"))
	})
	router.POST("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default post"))
	})
	// named parameters
	router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("user name:" + p.ByName("name")))
	})
	// catch-all parameters
	router.POST("/user/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("user name:" + p.ByName("name")))
	})

	// deel with different second-level domain names
	userRouter := httprouter.New()
	userRouter.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("sub1"))
	})

	dataRouter := httprouter.New()
	dataRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("sub2"))
	})

	//分别用于处理不同的二级域名
	hs := make(HostMap)
	hs["localhost:8888"] = router
	hs["sub1.localhost:8888"] = userRouter
	hs["sub2.localhost:8888"] = dataRouter

	http.ListenAndServe(":8888", hs)
}
