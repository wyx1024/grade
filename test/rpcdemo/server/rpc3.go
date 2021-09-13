package main

import (
	"go-growth/test/rpcdemo/etcd"
	pb "go-growth/test/rpcdemo/proto"
	"go-growth/test/rpcdemo/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	serve := grpc.NewServer()

	pb.RegisterTestServer(serve, &services.SayHello{})

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8557")
	if err != nil {
		panic(err)
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	reg, err := etcd.NewService(etcd.ServiceInfo{
		Name: "g.srv.mail",
		IP:   "127.0.0.1:8557", //grpc服务节点ip
	}, []string{"0.0.0.0:12379", "0.0.0.0:22379", "0.0.0.0:32379"}) // etcd的节点ip
	if err != nil {
		log.Fatal(err)
	}
	go reg.Start()

	serve.Serve(listen)
}
