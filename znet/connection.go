package znet

import (
	"errors"
	"fmt"
	"go-growth/utils"
	"go-growth/ziface"
	"io"
	"net"
)

type Connection struct {
	Conn   *net.TCPConn
	ConnID uint32

	isClose bool
	Router   ziface.IRouter
	ExitChan chan bool
}

func (c *Connection) StartRead() {
	fmt.Println("Start Connection Read ...ConID", c.ConnID)
	defer fmt.Println("Stop conn...ConnID", c.ConnID)
	defer c.Stop()
	for {
		buf := make([]byte, utils.GlobalObj.MaxPackageSize)
		num, err := c.Conn.Read(buf)
		if err != nil {
			if err == io.EOF{
				break
			}
			fmt.Println("recv read err :", err)
			continue
		}
		req := &Request{
			conn: c,
			data: buf[:num],
		}
		go func(r ziface.IRequeset) {
			c.Router.PreHeadler(r)
			c.Router.Headler(r)
			c.Router.PostHeadler(r)
		}(req)
	}
}

func (c *Connection) Start() {
	fmt.Println("Start connection ConnId...", c.ConnID)
	//开启一个读数据
	go c.StartRead()

	//开启一个写数据

}
func (c *Connection) Stop() {
	if c.isClose == true {
		return
	}
	c.isClose = true

	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetConn() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return errors.New("")
}

func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) ziface.IConnection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connId,
		isClose:  false,
		Router:   router,
		ExitChan: make(chan bool),
	}
	return c
}
