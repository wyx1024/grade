package znet

import (
	"fmt"
	"go-growth/utils"
	"go-growth/ziface"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int32
	Router    ziface.IRouter
}

func (s *Server) Start() {
	fmt.Println("[Zinx] Start APP name", utils.GlobalObj.Name, " IP:", utils.GlobalObj.Host, "Port:", utils.GlobalObj.TcpPort)
	fmt.Println("[Server] Start server Version", utils.GlobalObj.Version)
	go func() {
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
		fmt.Println("start tcp listener  "+s.Name+"success IP:"+s.IP+",Port:", s.Port)
		var connid uint32
		connid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				continue
			}
			fmt.Println("new conn success...ConnID...", connid)
			c := NewConnection(conn, connid, s.Router)

			go c.Start()

			connid++
		}
	}()
}
func (s *Server) Stop() {

}
func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObj.Name,
		IPVersion: "tcp",
		IP:        utils.GlobalObj.Host,
		Port:      utils.GlobalObj.TcpPort,
		Router:    nil,
	}
	return s
}
