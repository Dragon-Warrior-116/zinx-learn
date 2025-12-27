package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx-learn/ziface"
)

type server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func NewServer(name string) ziface.IServer {
	return &server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}

func CallBack(conn *net.TCPConn, data []byte, cnt int) error {
	//回显业务
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Printf("Write back err: %v\n", err)
		return errors.New("CallBack Write err")
	}
	return nil
}

func (s *server) Start() {
	go func() {
		// 1 初始化一个socket
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("ResolveTCPAddr err: %v\n", err)
			return
		}
		// 2 监听
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("ListenTCP err: %v\n", err)
			return
		}
		fmt.Printf("Server %s started, listening on %s\n", s.Name, listener.Addr().String())
		// 3 循环等待客户端的连接
		var cid uint32 = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("AcceptTCP err: %v\n", err)
				continue
			}
			// 4 创建连接对象
			c := NewConnection(conn, cid, CallBack)
			cid++
			// 5 启动连接
			go c.Start()
		}
	}()
}

func (s *server) Stop() {

}

func (s *server) Server() {
	s.Start()
}
