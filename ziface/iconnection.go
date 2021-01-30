package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConn() *net.TCPConn
	GetConnId() uint32
	GetRemoteAddr() net.Addr
	SendMsg(uint32, []byte) error
	SetPropety(string,interface{})
	GetPropety(string) (interface{}, error)
	RemovePropety(string)
}