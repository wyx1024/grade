package main

import (
	"fmt"
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
	for  {
		time.Sleep(time.Second*2)
		_, err =conn.Write([]byte("hello zinx v3.0"))
		if err != nil {
			if err != io.EOF {
				fmt.Println("server stop")
				break
			}
			fmt.Println("conn write err:", err)
			continue
		}

		buf := make([]byte, 512)
		num,err :=conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("server stop")
				break
			}
			fmt.Println("conn write err:", err)
			continue
		}
		fmt.Printf("%s\n", buf[:num])

	}
}
