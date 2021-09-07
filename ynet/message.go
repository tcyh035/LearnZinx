package ynet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte //消息内容
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(id uint32) {
	m.Id = id
}

func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}
