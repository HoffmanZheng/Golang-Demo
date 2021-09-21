package main

import (
	"net/http"
)

type Refer struct {
	handler http.Handler
	refer   string
}

func (referObj *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == referObj.refer {
		referObj.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func hello2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.shirdon.com",
	}
	http.HandleFunc("/hello", hello2)
	http.ListenAndServe(":80", referer)
}
