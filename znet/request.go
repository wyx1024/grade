package znet

import "go-growth/ziface"

type Request struct {
	conn ziface.IConnection
	Msg ziface.IMessage
}


func (req *Request)GetConnection() ziface.IConnection{
	return req.conn
}

func (req *Request)GetData() []byte{
	return req.Msg.GetData()
}

func (req *Request)GetMsgId() uint32  {
	return req.Msg.GetMsgID()
}