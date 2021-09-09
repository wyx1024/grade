package main

import (
	pb "go-growth/test/rpcdemo/proto"
	"go-growth/test/rpcdemo/services"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	serve := grpc.NewServer()

	pb.RegisterTestServer(serve, &services.SayHello{})

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8555")
	if err != nil {
		panic(err)
	}

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	serve.Serve(listen)
}
