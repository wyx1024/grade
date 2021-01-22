package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConn() *net.TCPConn
	GetConnId() uint32
	GetRemoteAddr() net.Addr
	Send(data []byte) error
}