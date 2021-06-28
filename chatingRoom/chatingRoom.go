package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn)  {
	fmt.Println("start handle success")
	// todo

}

func main() {
	// 创建服务
	listen, err := net.Listen("tcp", ":8079")
	if err != nil {
		fmt.Println("listen err:",err)
		return
	}
	fmt.Println("server start success")
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connect err:", err)
			return
		}
		// 建立连接
		fmt.Println("connect success")

		// 启动处理业务go程
		go handle(conn)
	}


}
