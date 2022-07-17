package main

import (
	"flag"
	"net"

	"log"
	stringService "stream_pb/stringService"

	stream_pb "github.com/HoffmanZheng/Golang-Demo/Go_Microservice_in_Action/chapter_7_remote_procedure_call/stream-pb"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	stringService := new(stringService.StringService)
	stream_pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}
