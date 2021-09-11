package main

import (
	"net/http"
)

type handler1 struct{}

func (h *handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here is handler 1!"))
}

type handler2 struct{}

func (h *handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here is handler 2!"))
}

func main() {
	handler1 := handler1{}
	handler2 := handler2{}
	server := http.Server{
		Addr:    ":80",
		Handler: nil,
	}
	http.Handle("/hi1", &handler1)
	http.Handle("/hi2", &handler2)
	server.ListenAndServe()
}
