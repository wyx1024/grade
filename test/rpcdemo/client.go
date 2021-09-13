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

func main() {
	r := etcd.NewResolver([]string{
		"0.0.0.0:12379",
		"0.0.0.0:22379",
		"0.0.0.0:32379",
	}, "g.srv.mail")
	resolver.Register(r)
	addr := fmt.Sprintf("%s:///%s", r.Scheme(), "g.srv.mail")
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTestClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "wuyanxiang"})
	time.Sleep(2*time.Second)
	fmt.Println(resp.String())
}



