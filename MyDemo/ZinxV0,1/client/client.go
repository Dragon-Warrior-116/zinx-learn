package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 1 连接服务器
	fmt.Printf("Client start...\n")
	conn, err := net.Dial("tcp", "0.0.0.0:8999")
	if err != nil {
		fmt.Printf("Dial err: %v\n", err)
		return
	}
	defer conn.Close()
	// 2 发送数据
	for {
		if _, err := conn.Write([]byte("Hello ZinxV0,1")); err != nil {
			fmt.Printf("Write err: %v\n", err)
			continue
		}
		// 读取服务器的响应
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read err: %v\n", err)
			continue
		}
		fmt.Printf("Server response: data:%s, cnt:%d\n", buf[:cnt], cnt)
		time.Sleep(3 * time.Second)
	}
}
