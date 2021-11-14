package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	Skill1, Skill2 string
}

type Programmer string // service obj

func (*Programmer) GetSkill(args *Args, skill *string) error {
	fmt.Printf("get rpc calls from client, args: %s %s \n", args.Skill1, args.Skill2)
	*skill = "Skill1: " + args.Skill1 + ", Skill2: " + args.Skill2
	return nil
}

func main() {
	programmer := new(Programmer)
	rpc.Register(programmer)

	l, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		fmt.Printf("error during net listen, err: %v \n", err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("error during connection accept, err: %v \n", err)
		}
		jsonrpc.ServeConn(conn)
	}
}
