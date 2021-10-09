package main

import (
	"context"
	"fmt"
	"go-growth/test/rpcdemo/etcd"
	pb "go-growth/test/rpcdemo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)


var addr string
var etcdEndPoint = []string{
	"0.0.0.0:12379",
	"0.0.0.0:22379",
	"0.0.0.0:32379",
}

func init() {
	r, err := etcd.NewResolver(etcdEndPoint, "g.srv.mail")

	if err != nil {
		panic(err)
	}
	resolver.Register(r)
	addr = fmt.Sprintf("%s:///%s", r.Scheme(), "g.srv.mail")
}

func main() {
	conn, err := grpc.DialContext(context.Background(), addr, grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewTestClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "wuyanxiang"})
	time.Sleep(2 * time.Second)
	fmt.Println(resp.String())
}
