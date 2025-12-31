package main

import (
	"zinx-learn/znet"
)

func main() {
	s := znet.NewServer("ZinxV0,1")
	s.Server()
	select {}
}