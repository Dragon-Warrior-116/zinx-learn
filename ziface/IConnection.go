package ziface

import "net"

// IConnection 定义了连接层对外暴露的接口
type IConnection interface {
	// Start 启动连接，开始读写处理
	Start()

	// Stop 停止连接，回收资源
	Stop()

	// GetConnID 获取当前连接的 ID
	GetConnID() uint32

	// GetTCPConnetion 获取原始的 TCP 连接对象（net.TCPConn）
	GetTCPConnetion() *net.TCPConn

	// RemoteAddr 获取远程客户端的地址信息
	RemoteAddr() net.Addr

	// SendMsg 发送消息（字节数组）给远程客户端
	SendMsg(data []byte) error
}

// HandleFunc 定义处理 TCP 连接读到的数据的函数签名
// 参数：TCP 连接对象、读到的数据、数据长度；返回值：处理过程中出现的错误
type HandleFunc func(*net.TCPConn, []byte, int) error
