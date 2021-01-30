package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"go-growth/utils"
	"go-growth/ziface"
)

type DataPack struct {
	
}

func NewDataPack() ziface.IDataPack  {
	return &DataPack{}
}
func (d DataPack) GetHeadLen() uint32 {
	//固定前八個字節存頭信息 len 和 msgid
	return 8
}

//封包
func (d DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//初始化一個緩存
	buf := bytes.NewBuffer([]byte{})
	//寫入數據長度
	if err := binary.Write(buf, binary.LittleEndian, msg.GetDataLen()); err != nil{
		fmt.Println("writer data len err :", err)
		return nil, err
	}
	//寫入msgID
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgID()); err != nil{
		fmt.Println("writer data msgid err :", err)
		return nil, err
	}

	//寫入內容
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil{
		fmt.Println("writer data msgid err :", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

//拆包
func (d DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	buf := bytes.NewBuffer(data)

	msg := &Messages{}

	if err := binary.Read(buf, binary.LittleEndian, &msg.Len); err != nil{
		fmt.Println("reader data len err :", err)
		return nil,err
	}

	if err := binary.Read(buf, binary.LittleEndian, &msg.MsgId); err != nil{
		fmt.Println("reader MsgId len err :", err)
		return nil,err
	}

	if utils.GlobalObj.MaxPackageSize > 0 && msg.Len > uint32(utils.GlobalObj.MaxPackageSize ){
		fmt.Println("err data size ")
		return nil, errors.New("err data size")
	}

	return msg, nil
}


