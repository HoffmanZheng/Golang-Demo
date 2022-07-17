package main

import (
	"context"
	"fmt"

	pb "github.com/HoffmanZheng/Golang-Demo/Go_Microservice_in_Action/chapter_7_remote_procedure_call/pb"
	"google.golang.org/grpc"
)

func main() {
	serviceAddress := "127.0.0.1:1234"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	bookClient := pb.NewStringServiceClient(conn)
	stringReq := &pb.StringRequest{A: "A", B: "B"}
	reply, _ := bookClient.Concat(context.Background(), stringReq)
	fmt.Printf("StringService Concat : %s concat %s = %s\n", stringReq.A, stringReq.B, reply.Ret)
}
