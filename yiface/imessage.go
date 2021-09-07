package yiface

type IMessage interface {
	GetMsgId() uint32
	GetDataLen() uint32
	GetData() []byte

	SetMsgID(uint32)
	SetMsgData([]byte)
	SetDataLen(uint32)
}
