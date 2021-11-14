package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("test_handle")
	fmt.Printf("cookie: %v, err: %v \n", cookie, err)

	setCookie := &http.Cookie{
		Name:   "test_handle",
		Value:  "gweoin2iehkljadhoiuh3",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/", // scope
	}

	http.SetCookie(w, setCookie)
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", setCookie)
	http.ListenAndServe(":8089", nil)
}
