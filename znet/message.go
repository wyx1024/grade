package znet

import "go-growth/ziface"

type Messages struct {
	Len   uint32
	MsgId uint32
	Data  []byte
}

func NewMessage(msgId uint32, data []byte) ziface.IMessage{
	return &Messages{
		Len:   uint32(len(data)),
		MsgId: msgId,
		Data:  data,
	}
}

func (m *Messages) GetDataLen() uint32 {
	return m.Len
}

func (m *Messages) GetMsgID() uint32 {
	return m.MsgId
}

func (m *Messages) GetData() []byte {
	return m.Data
}

func (m *Messages) SetDataLen(len uint32) {
	m.Len = len
}

func (m *Messages) SetMsgID(msgid uint32) {
	m.MsgId = msgid
}

func (m *Messages) SetData(data []byte) {
	m.Data = data
}
