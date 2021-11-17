package main

import (
	"context"
	"fmt"

	pb "github.com/NervousOrange/Golang-demo/src/chapter_5_socket_rpc/protobuf"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8091", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error during grpc dial, err: %v \n", err)
	}
	defer conn.Close()

	client := pb.NewProgrammerServiceClient(conn)

	// 构造 rpc 调用参数
	req := new(pb.Request)
	req.Name = "Hoffman"
	resp, err := client.GetProgrammerInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("error during client rpc getProgrammerInfo, err: %v \n", err)
	}
	fmt.Printf("response from rpc server: %v \n", resp)

}
