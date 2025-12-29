package ziface

// IRequest 定义了客户端请求消息的抽象接口
type IRequest interface {
	// GetConnection 获取当前请求对应的连接对象
	GetConnection() IConnection
	// GetData 获取当前请求携带的原始字节数据
	GetData() []byte
}
