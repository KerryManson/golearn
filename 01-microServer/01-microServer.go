package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type World struct {
}

func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + "你好"
	fmt.Println(*resp)
	return nil
}

func main() {
	//err := rpc.RegisterName("hello", new(World))
	//if err != nil {
	//	fmt.Println("注册helloRPC服务失败:",err)
	//	return
	//}
	RegisterService(new(World))

	listener, err := net.Listen("tcp", ":8023")
	if err != nil {
		fmt.Println("start Listening failed:", err)
		return
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("connect failed:", err)
		return
	}
	defer conn.Close()
	fmt.Println("has bound service")
	rpc.ServeConn(conn)

}
