package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

type Args struct {
	Skill1, Skill2 string
}

func main() {
	fmt.Println("Client start...")
	client, err := jsonrpc.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Printf("error during jsonrpc dial, err: %v \n", err)
	}
	send := Args{"Golang", "Java"}
	var result string
	fmt.Printf("args in Send, skill1: [%s], skill2: [%s] \n", send.Skill1, send.Skill2)
	err2 := client.Call("Programmer.GetSkill", send, &result)
	if err2 != nil {
		fmt.Printf("error during rpc call, err: %v \n", err)
	}
	fmt.Printf("get result from rpcServer: %s \n", result)
}
