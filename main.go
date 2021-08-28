package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"yinx/yiface"
	"yinx/ynet"
)

type MyRouter struct {
	ynet.BaseRouter
}

func (router *MyRouter) Handle(request yiface.IRequest) {

	fmt.Println("[My Router] Handle")
	conn := request.GetConnection()
	s := strings.ToUpper(string(request.GetData()))
	_, err := conn.GetTCPConnection().Write([]byte(s))
	if err != nil {
		fmt.Println("write back buf err:", err)
	}
}

func CallbackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallbackToClient")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err:", err)
		return errors.New("CallbackToClient error")
	}

	return nil
}

func main() {
	s := ynet.NewServer("Hello")
	s.AddRouter(&MyRouter{})
	s.Serve()
}
