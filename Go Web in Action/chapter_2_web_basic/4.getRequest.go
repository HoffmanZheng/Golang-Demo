package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.DefaultClient.Get("http://www.baidu.com")
	if err != nil {
		fmt.Print("error during getting request: \n", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("error during reading response: \n", err)
	}
	fmt.Println(string(bytes))
}
