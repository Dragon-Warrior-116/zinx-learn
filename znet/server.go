package znet

import (
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
		// 循环等待客户端的连接
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("AcceptTCP err: %v\n", err)
				continue
			}
			// 处理回显
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("Read err: %v\n", err)
						continue
					}
					fmt.Printf("Client %s send: %s, cnt:%d\n", conn.RemoteAddr().String(), buf[:cnt], cnt)
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Printf("Write err: %v\n", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *server) Stop() {

}

func (s *server) Server() {
	s.Start()
}
