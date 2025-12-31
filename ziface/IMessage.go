package ziface

type IMessage interface {
	// GetID 获取消息ID
	GetID() uint32

	// GetDataLen 获取消息内容长度
	GetDataLen() uint32

	// GetData 获取消息内容
	GetData() []byte

	// SetID 设置消息ID
	SetID(uint32)

	// SetDataLen 设置消息内容长度
	SetDataLen(uint32)

	// SetData 设置消息内容
	SetData([]byte)
}
