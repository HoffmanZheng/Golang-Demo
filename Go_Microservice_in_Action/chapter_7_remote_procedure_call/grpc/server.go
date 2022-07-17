package main

import (
	"flag"
	"net"

	string_service "grpc/string-service"

	"github.com/HoffmanZheng/Golang-Demo/Go_Microservice_in_Action/chapter_7_remote_procedure_call/pb"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	stringService := new(string_service.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}
