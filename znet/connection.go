package znet

import (
	"errors"
	"fmt"
	"go-growth/utils"
	"go-growth/ziface"
	"io"
	"net"
	"sync"
)

type Connection struct {
	TcpServer ziface.IServer
	Conn   *net.TCPConn
	ConnID uint32

	isClose    bool
	MsgHandler ziface.IMsgHandle
	ExitChan   chan bool
	//读写通信
	msgChan chan []byte
	Propety map[string]interface{}
	PropetyLock sync.RWMutex
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connId uint32, msghandler ziface.IMsgHandle) ziface.IConnection {
	c := &Connection{
		TcpServer:server,
		Conn:       conn,
		ConnID:     connId,
		isClose:    false,
		MsgHandler: msghandler,
		ExitChan:   make(chan bool),
		msgChan:    make(chan []byte, 1),
		Propety: make(map[string]interface{}),
	}
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Connection] Start Reader ...ConID", c.ConnID)
	defer fmt.Println(c.GetRemoteAddr().String(), "[Conn Reader exit]")
	defer fmt.Println("[Connection] Stop StartRead...ConnID", c.ConnID)
	defer c.Stop()
	for {
		//改成包裝好的格式拆包读取数据
		dp := NewDataPack()
		dataHead := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetConn(), dataHead)
		if err != nil {
			fmt.Println("read head data err ", err)
			break
		}
		msg, err := dp.UnPack(dataHead)
		if err != nil {
			fmt.Println("unpack head dat err :", err)
			return
		}
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			_, err = io.ReadFull(c.GetConn(), data)
			if err != nil {
				fmt.Println("read data err:", err)
				return
			}
			msg.SetData(data)
		}

		req := &Request{
			conn: c,
			Msg:  msg,
		}
		if utils.GlobalObj.WorkerPoolSize > 0 {
			go c.MsgHandler.SendMsgToQueue(req)
		}else {
			go c.MsgHandler.DoMsgHandle(req)
		}
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("[Connection] Start Writer ...ConID", c.ConnID)
	defer fmt.Println(c.GetRemoteAddr().String(), "[Conn Writer exit]")
	defer fmt.Println("[Connection] Stop StartWriter...ConnID", c.ConnID)
	for {
		select {
		case data:= <- c.msgChan:
			if _, err := c.Conn.Write(data);err != nil{
				fmt.Println("Send data err", err)
				return
			}
		case <- c.ExitChan:
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("[Connection] Start ConnId...", c.ConnID)
	//开启一个读数据
	go c.StartReader()

	//开启一个写数据
	go c.StartWriter()

	//调用hook函数
	c.TcpServer.CallOnConnStart(c)

}
func (c *Connection) Stop() {
	if c.isClose == true {
		return
	}
	c.isClose = true

	c.TcpServer.CallOnConnStop(c)

	c.Conn.Close()
	c.TcpServer.GetConnMgr().Remove(c)
	c.ExitChan <- true
	close(c.ExitChan)
	close(c.msgChan)
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClose {
		fmt.Println("[Connection] close ")
		return errors.New("[Connection] close")
	}

	dp := NewDataPack()

	binaryData, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("pack err ", err)
		return err
	}

	c.msgChan <- binaryData
	return nil
}


func (c *Connection)SetPropety(key string,value interface{}){
	c.PropetyLock.Lock()
	defer c.PropetyLock.Unlock()
	c.Propety[key] = value
}

func (c *Connection)GetPropety(key string) (interface{}, error){
	c.PropetyLock.RLock()
	defer c.PropetyLock.RUnlock()
	if value ,ok:=  c.Propety[key];ok{
		return value,nil
	}
	return nil, errors.New("[Connection] GetPropety Not Found Value")
}

func (c *Connection)RemovePropety(key string){
	delete(c.Propety, key)
}


