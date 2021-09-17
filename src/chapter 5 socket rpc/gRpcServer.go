package main

import (
	"context"

	pb "github.com/NervousOrange/Golang-demo/src/chapter_5_socket_rpc"
)

type ProgrammerServiceServer struct{}

func (p *ProgrammerServiceServer) GetProgrammerInfo(ctx context.Context,
	req *pb.Request) (resp *pb.Response, err error) {
	name = req.name
	if name == "Hoffman" {
		resp = &pb.Response{
			Uid:      6,
			Username: name,
			Job:      "CTO",
			GoodAt:   []String{"Go", "Java", "Sql"},
		}
	}
	err = nil
	return
}

func main() {

}
