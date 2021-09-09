package services

import (
	"context"
	pb "go-growth/test/rpcdemo/proto"
)

type SayHello struct {
}

func (say SayHello) SayHello(ctx context.Context ,req *pb.HelloRequest) (resp *pb.HelloResponse,err error) {
	resp = &pb.HelloResponse{}
	name := req.GetName()
	resp.Msg = "hello rpc "+name
	return resp, nil
}

