package main

import (
	"fmt"
	"yinx/yiface"
	"yinx/ynet"
)

type MyRouter struct {
	ynet.BaseRouter
}

func (router *MyRouter) Handle(request yiface.IRequest) {
	fmt.Println("[My Router] Handle")
	conn := request.GetConnection()
	received := string(request.GetMessage().GetData())
	fmt.Println("len is ", len(received))
	_, err := conn.GetTCPConnection().Write([]byte(received))
	if err != nil {
		fmt.Println("write back buf err:", err)
	}
}

func main() {
	s := ynet.NewServer("Hello")
	s.AddRouter(&MyRouter{})
	s.Serve()
}
