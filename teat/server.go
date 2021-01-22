package main

import (
	"fmt"
	"go-growth/ziface"
	"go-growth/znet"
)

type TeatAPI struct {
	znet.BaseRouter
}
func(r *TeatAPI) PreHeadler(req ziface.IRequeset) {
	fmt.Println("pre headler ...ConnID", req.GetConnection().GetConnId())
}
func(r *TeatAPI)Headler(req ziface.IRequeset){
	fmt.Println("headler ...ConnID", req.GetConnection().GetConnId())
	req.GetConnection().GetConn().Write(req.GetData())
}
func(r *TeatAPI)PostHeadler(req ziface.IRequeset){
	fmt.Println("post headler ...ConnID", req.GetConnection().GetConnId())
}


func main() {
	s := znet.NewServer("zinx3.0")

	s.AddRouter(new(TeatAPI))

	s.Serve()
}
