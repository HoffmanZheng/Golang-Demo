package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8088", Handler: http.HandlerFunc(myHandler2)}
	log.Printf("Serving on https://0.0.0.0:8088")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	w.Write([]byte("Hello this is a HTTP 2 message!"))
}
