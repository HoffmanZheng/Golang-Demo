package main

import (
	"net/http"
)

func handleFunc1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here is handler1"))
}

func handleFunc2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here is handler2"))
}

func main() {
	myServeMux := http.NewServeMux()
	myServeMux.HandleFunc("/hi1", handleFunc1)
	myServeMux.HandleFunc("/hi2", handleFunc2)
	myServer := http.Server{
		Handler: myServeMux,
	}
	myServer.ListenAndServe()
}
