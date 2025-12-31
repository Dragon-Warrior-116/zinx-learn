package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx-learn/utils"
	"zinx-learn/ziface"
)

type Datapack struct{}

func NewDatapack() ziface.IDatapack {
	return &Datapack{}
}

func (dp *Datapack) GetHeadLen() uint32 {
	return 8
}

func (dp *Datapack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetID()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (dp *Datapack) Unpack(data []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewBuffer(data)
	msg := &Message{}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//判断dataLen是否超出了我们允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data recv")
	}
	return msg, nil
}
