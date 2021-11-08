package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.shirdon.com/comment/delete"
	payload := strings.NewReader("{\"userId\":1,\"articleId\":1,\"comment\":\"new comment\"}")
	req, err := http.NewRequest("DELETE", url, payload)
	if err != nil {
		fmt.Print("error during newRequest: \n", err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print("error during delete: \n", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Print(bytes)
}
