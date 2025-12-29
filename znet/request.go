// Request 实现 ziface.IRequest 接口，封装客户端连接及其请求数据
package znet

import "zinx-learn/ziface"

// Request 结构体保存一次客户端请求的完整信息
type Request struct {
	conn ziface.IConnection // 客户端连接
	data []byte             // 请求数据
}

// GetConnection 返回当前请求对应的客户端连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 返回当前请求携带的原始数据
func (r *Request) GetData() []byte {
	return r.data
}
