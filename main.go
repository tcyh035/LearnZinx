package main

import "yinx/ynet"

func main() {
	s := ynet.NewServer("Hello")
	s.Serve()
}
