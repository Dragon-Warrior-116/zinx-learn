package znet

import (
	"fmt"
	"net"
	"zinx-learn/utils"
	"zinx-learn/ziface"
)

// Connection 代表一个客户端连接
type Connection struct {
	Conn *net.TCPConn // 底层TCP连接

	ConnID uint32 // 连接ID

	IsClose bool // 连接是否已关闭

	ExitChan chan bool // 用于通知连接退出的通道

	Router ziface.IRouter // 路由
}

// NewConnection 创建一个新的连接对象
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		IsClose:  false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
	return c
}

// Start 启动连接，主要是启动读取数据的goroutine
func (c *Connection) Start() {
	fmt.Printf("Connection %d started\n", c.ConnID)
	go c.StartReader() // 在单独的goroutine中启动读取数据
	//TODO 启动写入数据的goroutine
}

// StartReader 启动读取客户端数据的goroutine，持续读取客户端发送的数据
func (c *Connection) StartReader() {
	fmt.Printf("Reader Goroutine is Runnig, Connection %d\n", c.ConnID)
	defer c.Stop()
	for {
		//读客户端的数据到buf
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("Read err: %v\n", err)
			continue
		}
		req := &Request{
			conn: c,
			data: buf,
		}
		go func(req *Request) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(req)
	}
}

// Stop 停止连接，关闭相关资源
func (c *Connection) Stop() {
	if c.IsClose {
		return
	}
	c.IsClose = true
	// 关闭连接
	c.Conn.Close()
	close(c.ExitChan)
}

// GetConnID 获取连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// GetTCPConnetion 获取底层TCP连接
func (c *Connection) GetTCPConnetion() *net.TCPConn {
	return c.Conn
}

// RemoteAddr 获取远程客户端地址
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 发送消息到客户端
func (c *Connection) SendMsg(data []byte) error {
	return nil
}
