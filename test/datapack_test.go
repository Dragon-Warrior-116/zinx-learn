package test

import (
	"fmt"
	"net"
	"testing"
	"time"
	"zinx-learn/znet"
)

func TestDatapack_Pack_Unpack(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:8999")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Fatal(err)
			}
			defer conn.Close()
			go func(conn net.Conn) {
				dp := znet.NewDatapack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					if _, err := conn.Read(headData); err != nil {
						t.Fatal(err)
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						t.Fatal(err)
						break
					}
					if msgHead.GetDataLen() > 0 {
						msg := &znet.Message{}
						msg.Data = make([]byte, msgHead.GetDataLen())
						if _, err := conn.Read(msg.Data); err != nil {
							t.Fatal(err)
							break
						}
						fmt.Printf("msgID: %d, dataLen: %d, data: %s\n", msgHead.GetID(), msgHead.GetDataLen(), string(msg.Data))
					}
				}
			}(conn)
		}
	}()
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		t.Fatal(err)
		return
	}
	dp := znet.NewDatapack()
	msg := &znet.Message{
		ID:      1,
		DataLen: 5,
		Data:    []byte("hello"),
	}
	packedData1, err := dp.Pack(msg)
	if err != nil {
		t.Fatal(err)
		return
	}
	msg2 := &znet.Message{
		ID:      2,
		DataLen: 5,
		Data:    []byte("world"),
	}
	packedData2, err := dp.Pack(msg2)
	if err != nil {
		t.Fatal(err)
		return
	}
	data := append(packedData1, packedData2...)
	for {
		conn.Write(data)
		time.Sleep(3 * time.Second)
	}
}
