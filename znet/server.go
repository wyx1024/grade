package znet

import (
	"fmt"
	"go-growth/utils"
	"go-growth/ziface"
	"net"
)

type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int32
	MsgHandler  ziface.IMsgHandle
	connMgr     ziface.IConnManager
	OnConnStart func(ziface.IConnection)
	OnConnStop  func(ziface.IConnection)
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObj.Name,
		IPVersion:  "tcp",
		IP:         utils.GlobalObj.Host,
		Port:       utils.GlobalObj.TcpPort,
		MsgHandler: NewMsgHandle(),
		connMgr:    NewConnManager(),
	}
	return s
}

func (s *Server) Start() {
	fmt.Println("[Zinx] Start APPName", utils.GlobalObj.Name, ",IP:", utils.GlobalObj.Host, ",Port:", utils.GlobalObj.TcpPort)
	fmt.Println("[Zinx] Start server Version", utils.GlobalObj.Version,
		",MaxConn", utils.GlobalObj.MaxConn,
		",MaxPackageSize", utils.GlobalObj.MaxPackageSize,
		",WorkerPoolSize", utils.GlobalObj.WorkerPoolSize)
	go func() {
		s.MsgHandler.StartWorkerPool()
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Start server", s.Name, "resolve tcp addr err:", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Start server", s.Name, " tcp listener err:", err)
			return
		}
		fmt.Println("[Server] ListenTCP Success:IP:"+s.IP+",Port:", s.Port)
		var connid uint32
		connid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				continue
			}
			if s.connMgr.GetLen() == utils.GlobalObj.MaxConn {
				fmt.Println("Now Conn Len=", s.connMgr.GetLen(), ",Max ConnSize", utils.GlobalObj.MaxConn)
				conn.Close()
				continue
			}
			fmt.Println("[Server] AcceptTCP success...ConnID...", connid)
			c := NewConnection(s, conn, connid, s.MsgHandler)

			go c.Start()

			connid++
		}
	}()
}
func (s *Server) Stop() {
	fmt.Println("[Server] Stop Server")
	s.connMgr.ClearConn()
}
func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(msgid uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgid, router)
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.connMgr
}

func (s *Server) SetOnConnStart(hookfunc func(conn ziface.IConnection)) {
	s.OnConnStart = hookfunc
}

func (s *Server) SetOnConnStop(hookfunc func(conn ziface.IConnection)) {
	s.OnConnStop = hookfunc
}
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil{
		fmt.Println("[Server] Call OnConnStart")
		s.OnConnStart(conn)
	}
}
func (s *Server) CallOnConnStop(conn ziface.IConnection)  {
	if s.OnConnStop != nil{
		fmt.Println("[Server] Call OnConnStop")
		s.OnConnStop(conn)
	}
}
