package ynet

import (
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

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept failed, err:", err)
				continue
			}

			go func() {
				fmt.Println("Accept! from", conn.RemoteAddr().String())
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Read failed, err:", err)
						continue
					}

					fmt.Println("接受到了！", string(buf[:n]))

					_, err = conn.Write(buf[:n])
					if err != nil {
						fmt.Println("Write failed, err", err)
						continue
					}
				}
			}()
		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	// 其他操作

	select {}
}

func NewServer(name string) yiface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
	}

	return s
}
