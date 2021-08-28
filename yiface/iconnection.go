package yiface

import "net"

type IConnection interface {
	// 启动连接，让当前连接开始工作
	Start()

	// 结束当前连接的工作
	Stop()

	// 获取当前socket conn
	GetTCPConnection() *net.TCPConn

	// 获取当前连接模块的id
	GetConnID() uint32

	// 获取远程客户端的TCP状态
	RemoteAddr() net.Addr

	// // 发送数据
	// Send(data []byte) error
}