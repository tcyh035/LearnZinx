package ynet

import (
	"fmt"
	"net"
	"yinx/utils"
	"yinx/yiface"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	Router yiface.IRouter

	// 告知当前连接已经退出/停止 channel
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, router yiface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}

	return c
}

// 读客户端信息的业务
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running")
	defer fmt.Println("connID =", c.ConnID, "Reader exits, remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中，目前最大512字节
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			return
		}

		request := Request{
			conn: c,
			data: buf[:cnt],
		}

		go func(request yiface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&request)

	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start() ... ConnID =", c.ConnID)
	go c.StartReader()
	// 启动从当前连接读数据的业务
	// TODO 启动从当前写数据的业务, 现在都写一起了
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID =", c.ConnID)

	if c.isClosed {
		return
	}

	c.isClosed = true

	// 关闭 socket
	c.Conn.Close()

	// 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
