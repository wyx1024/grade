package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDatapack(t *testing.T) {
	//server端
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:7777")
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil{
		fmt.Println(err)
	}
	go func() {
		for  {
			c ,err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP", err)
				continue
			}
			go func(conn net.Conn) {
				fmt.Println("---------conn success----------")
				//讀取頭信息
				dp := NewDataPack()
				for {
					haadData := make([]byte, dp.GetHeadLen())
					_, err = io.ReadFull(conn, haadData)
					if err != nil {
						fmt.Println("ReadFull:",err)
					}
					HeadMsg, err  := dp.UnPack(haadData)
					if err != nil {
						fmt.Println("ReadFull:",err)
					}
					if HeadMsg.GetDataLen() > 0 {
						msg := HeadMsg.(*Messages)
						msg.Data = make([]byte, msg.GetDataLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("ReadFull:",err)
						}
						fmt.Println("==> Recv Msg: ID=", msg.MsgId, ", len=", msg.Len, ", data=", string(msg.Data))
					}
				}
			}(c)
		}
	}()

	//客戶端
	go func() {
		conn, err := net.Dial("tcp", "127.0.0.1:7777")
		if err != nil {
			fmt.Println(err)
		}
		dp := NewDataPack()

		msg1 := &Messages{
			Len:   5,
			MsgId: 1,
			Data:  []byte{'h', 'e', 'l', 'l', 'o'},
		}
		data,err := dp.Pack(msg1)
		if err != nil {
			fmt.Println(err)
		}

		msg2 := &Messages{
			Len:   7,
			MsgId: 2,
			Data:   []byte{'w', 'o', 'r', 'l', 'd','!', '!'},
		}
		data2 ,err := dp.Pack(msg2)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, data2...)
		fmt.Println("---------writer success--------")
		conn.Write(data)
	}()

	//客户端阻塞
	select {
	case <-time.After(2*time.Second):
		return
	}
}