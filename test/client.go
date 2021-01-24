package main

import (
	"fmt"
	"go-growth/znet"
	"io"
	"net"
	"time"
)

func main() {
	conn, err:=net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("dial tcp err: ", err)
		return
	}
	dp := znet.NewDataPack()
	var msgID uint32
	msgID = 2
	for  {
		time.Sleep(time.Second)
		binaryData, err := dp.Pack(znet.NewMessage(msgID,[]byte("hello zinx v6.0")))
		_, err =conn.Write(binaryData)
		if err != nil {
			if err != io.EOF {
				fmt.Println("server stop")
				break
			}
			fmt.Println("conn write err:", err)
			continue
		}

		dataHead := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, dataHead)
		if err != nil {
			if err != io.EOF {
				fmt.Println("server stop")
				break
			}
			fmt.Println("read head data err ", err)
			break
		}
		msg, err := dp.UnPack(dataHead)
		if err != nil {

			fmt.Println("unpack head dat err :", err)
			continue
		}
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			_, err = io.ReadFull(conn, data)
			if err != nil {
				fmt.Println("read data err:", err)
				return
			}
			msg.SetData(data)
		}
		fmt.Println("msg", msg.GetMsgID(),msg.GetDataLen(), string(msg.GetData()))
		if msgID == 2 {
			msgID = 1
		}else {
			msgID = 2
		}
	}
}
