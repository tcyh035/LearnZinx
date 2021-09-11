package ynet

import (
	"fmt"
	"io"
	"net"
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
		dp := NewDataPack()
		// 第一次读，把数据头读出来
		dataHead := make([]byte, dp.GetHeadLen())
		// ReadFull如果没读满的话会抛异常，故不需要使用返回的长度
		_, err := io.ReadFull(c.Conn, dataHead)
		if err != nil {
			fmt.Println("read data head failed, err", err)
			break
		}

		// 此时解析数据头，生成message，message里只有长度，没有数据本体
		msg, err := dp.Unpack(dataHead)
		if err != nil {
			fmt.Println("unpack message error", err)
		}

		// 第二次读，把数据体读出来
		var dataBody []byte
		if msg.GetDataLen() > 0 {
			dataBody = make([]byte, msg.GetDataLen())
			_, err = io.ReadFull(c.Conn, dataBody)
			if err != nil {
				fmt.Println("read data body failed, err", err)
				break
			}
		}

		msg.SetMsgData(dataBody)

		request := Request{
			conn:    c,
			message: msg,
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
