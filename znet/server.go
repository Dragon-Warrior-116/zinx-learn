package znet

import (
	"fmt"
	"net"
	"zinx-learn/utils"
	"zinx-learn/ziface"
)

type server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    ziface.IRouter
}

func NewServer(name string) ziface.IServer {
	return &server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
}

func (s *server) Start() {
	go func() {
		// 1 初始化一个socket
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", utils.GlobalObject.Host, utils.GlobalObject.TcpPort))
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
			c := NewConnection(conn, cid, s.Router)
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

func (s *server) AddRouter(router ziface.IRouter) {
	// 1 注册路由
	s.Router = router
	fmt.Printf("AddRouter: %v\n", router)
}
