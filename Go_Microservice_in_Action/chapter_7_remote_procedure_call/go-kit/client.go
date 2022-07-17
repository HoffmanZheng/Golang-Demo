package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	service "go_kit/stringService"

	pb "github.com/HoffmanZheng/Golang-Demo/Go_Microservice_in_Action/chapter_7_remote_procedure_call/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		fmt.Println("gRPC dial err:", err)
	}
	defer conn.Close()

	svr := NewStringClient(conn)
	result, err := svr.Concat(ctx, "A", "B")
	if err != nil {
		fmt.Println("Check error", err.Error())
	}

	fmt.Println("result=", result)
}

func NewStringClient(conn *grpc.ClientConn) service.Service {

	var ep = grpctransport.NewClient(conn,
		"pb.StringService",
		"Concat",
		DecodeStringRequest,
		EncodeStringResponse,
		pb.StringResponse{},
	).Endpoint()

	userEp := service.StringEndpoints{
		StringEndpoint: ep,
	}
	return userEp
}

func DecodeStringRequest(ctx context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

func EncodeStringResponse(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}
