package ynet

import (
	"errors"
	"fmt"
	"net"
	"yinx/yiface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func CallbackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallbackToClient")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err:", err)
		return errors.New("CallbackToClient error")
	}

	return nil
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

			dealConn := NewConnection(conn, cid, CallbackToClient)
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

func NewServer(name string) yiface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}
