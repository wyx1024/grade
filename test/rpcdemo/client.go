package main

import (
	"context"
	"fmt"
	pb "go-growth/test/rpcdemo/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn , err :=grpc.Dial("127.0.0.1:8555", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewTestClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "wuyanxiang"})
	time.Sleep(2*time.Second)
	fmt.Println(resp.String())
}



