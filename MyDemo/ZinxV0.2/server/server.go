package main

import (
	"fmt"
	"zinx-learn/ziface"
	"zinx-learn/znet"
)

type PingTest struct {
	znet.BaseRouter
}

func (p *PingTest) PreHandle(request ziface.IRequest) {
	fmt.Println("Call PingTest PreHandle")
	conn := request.GetConnection()
	conn.GetTCPConnetion().Write([]byte("PreHandle\n"))
}

func (p *PingTest) Handle(request ziface.IRequest) {
	fmt.Println("Call PingTest Handle")
	conn := request.GetConnection()
	conn.GetTCPConnetion().Write([]byte("Handle\n"))
}

func (p *PingTest) PostHandle(request ziface.IRequest) {
	fmt.Println("Call PingTest PostHandle")
	// 1 从request中获取连接
	conn := request.GetConnection()
	conn.GetTCPConnetion().Write([]byte("PostHandle\n"))
}

func main() {
	s := znet.NewServer("ZinxV0.2")
	s.AddRouter(&PingTest{})
	s.Server()
	select {}
}
