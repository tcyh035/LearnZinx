package yiface

type IRequest interface {
	GetConnection() IConnection

	GetMessage() IMessage

	GetMsgID() uint32
}
