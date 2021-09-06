package ynet

import (
	"fmt"
	"net"
	"yinx/utils"
	"yinx/yiface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    yiface.IRouter
}

func (s *Server) Start() {
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Resolve TCP Address Failed, err:", err)
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen failed, err:", err)
		}

		var cid uint32 = 0

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept failed, err:", err)
				continue
			}

			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	fmt.Println("[Server]", s.Name, "start to serve")
	s.Start()

	// 其他操作

	select {}
}

func (s *Server) AddRouter(router yiface.IRouter) {
	s.Router = router
	fmt.Println("Add router!")
}

func NewServer(name string) yiface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}

	return s
}
