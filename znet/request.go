package znet

import "go-growth/ziface"

type Request struct {
	conn ziface.IConnection
	data []byte
}


func (req *Request)GetConnection() ziface.IConnection{
	return req.conn
}

func (req *Request)GetData() []byte{
	return req.data
}