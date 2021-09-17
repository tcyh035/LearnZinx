package ynet

import "yinx/yiface"

type Request struct {
	// 已经和客户端建立好的连接
	conn yiface.IConnection

	// 数据
	message yiface.IMessage
}

func (r *Request) GetConnection() yiface.IConnection {
	return r.conn
}

func (r *Request) GetMessage() yiface.IMessage {
	return r.message
}

func (r *Request) GetMsgID() uint32 {
	return r.message.GetMsgId()
}
