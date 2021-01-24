package main

import (
	"fmt"
	"go-growth/ziface"
	"go-growth/znet"
	"time"
)

type TeatAPI struct {
	znet.BaseRouter
}
func(r *TeatAPI)Headler(req ziface.IRequeset){
	fmt.Println("[API] Headler...ConnID", req.GetConnection().GetConnId(), ",MsgID=",req.GetMsgId())
	req.GetConnection().SendMsg(req.GetMsgId(), req.GetData())
}

type TeatAPI2 struct {
	znet.BaseRouter
}

func(r *TeatAPI2)Headler(req ziface.IRequeset){
	fmt.Println("[API] headler...ConnID", req.GetConnection().GetConnId(),",MsgID=", req.GetMsgId())
	req.GetConnection().SendMsg(req.GetMsgId(), req.GetData())
}

func DnConnStart(conn ziface.IConnection)  {
	fmt.Println("DnConnStart")
	conn.SendMsg(401, []byte("DnConnStart"))
	conn.SetPropety("Name","wuyanxiang")
	conn.SetPropety("Time",time.Now().Format("2006-01-02 15:04:05"))
	conn.SetPropety("Version","v1.0")
}
func DoConnStop(conn ziface.IConnection)  {
	fmt.Println("DoConnStop")
	if value, err:=conn.GetPropety("Name");err == nil{
		fmt.Println("Name", value)
	}
	if value, err:=conn.GetPropety("Time");err == nil{
		fmt.Println("Time", value)
	}
	if value, err:=conn.GetPropety("Version");err == nil{
		fmt.Println("Version", value)
	}
	conn.SendMsg(402, []byte("DoConnStop"))
}


func main() {
	s := znet.NewServer("zinx")
	s.SetOnConnStart(DnConnStart)
	s.SetOnConnStop(DoConnStop)

	s.AddRouter(1,new(TeatAPI))
	s.AddRouter(2,new(TeatAPI2))
	s.Serve()
}
