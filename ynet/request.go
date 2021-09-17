package ynet

import "yinx/yiface"

type Request struct {
	// 已经和客户端建立好的连接
	conn yiface.IConnection

	// 数据
	msg yiface.IMessage
}

func (r *Request) GetConnection() yiface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
