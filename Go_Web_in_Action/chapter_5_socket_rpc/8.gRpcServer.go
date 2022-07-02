package main

import (
	"fmt"
	"net"

	pb "github.com/HoffmanZheng/Golang-Demo/Go_Web_in_Action/chapter_5_socket_rpc/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ProgrammerServiceServer struct{}

func (p *ProgrammerServiceServer) GetProgrammerInfo(ctx context.Context,
	req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "Hoffman" {
		resp = &pb.Response{
			Uid:      6,
			Username: name,
			Job:      "CTO",
			GoodAt:   []string{"Go", "Java", "Sql"},
		}
	}
	err = nil
	return
}

func main() {
	l, err := net.Listen("tcp", "localhost:8091")
	if err != nil {
		fmt.Printf("error during net listenning, err: %v \n", err)
	}
	s := grpc.NewServer()
	// 第二个参数为接口类型的变量，需要取地址传参
	pb.RegisterProgrammerServiceServer(s, &ProgrammerServiceServer{})
	s.Serve(l)
}
