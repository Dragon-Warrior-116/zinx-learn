package ziface

type IDatapack interface {
	GetHeadLen() uint32
	// Pack 打包消息，将消息转换为字节流
	Pack(msg IMessage) ([]byte, error)

	// Unpack 解包消息，将字节流转换为消息
	Unpack([]byte) (IMessage, error)
}
