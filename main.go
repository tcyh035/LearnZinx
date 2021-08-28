package main

import (
	"fmt"
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

func main() {
	s := ynet.NewServer("Hello")
	s.AddRouter(&MyRouter{})
	s.Serve()
}
