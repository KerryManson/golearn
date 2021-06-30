package main

import (
	"fmt"
	"net"
	"time"
)

/*
1.有很多用户
2. 有个user结构，主要包含msg管道
3. 有个进行全局广播的管道
4. 单独go程服务器获取用户内容后 发送到全局用户中
5.全局的map管理所有的user
*/

type User struct {
	id   string
	name string
	msg  chan string
}

var AllUser = make(map[string]User)

var message = make(chan string)

func broadcast() {
	/*
		向全部用户广播, 全局唯一go程
	*/
	fmt.Println("broadCast server go rounte start suss")
	for {
		info := <-message
		fmt.Println("broadcast info", info)

		for _, user := range AllUser {
			user.msg <- info
		}
	}
}

func handle(conn net.Conn) {
	fmt.Println("start handle success")
	ClientInfo := conn.RemoteAddr().String()
	// fmt.Println(ClientInfo)
	newUser := User{
		id:   ClientInfo,
		name: ClientInfo,
		msg:  make(chan string, 10),
	}
	defer func() {
		fmt.Println("用户" + newUser.name + "下线")
		delete(AllUser, ClientInfo)
	}()
	go func() {
		// 向AllUser 中添加 key
		AllUser[newUser.id] = newUser
		fmt.Println(AllUser)
		// 讲登录数据同步到服务器
		sprintf := fmt.Sprintf("%s:%s ====> ONLINE\r\n", newUser.name, newUser.id)
		message <- sprintf

		messageInfo := <-newUser.msg
		fmt.Println("messageInfo:", messageInfo)
		conn.Write([]byte(messageInfo))
	}()

	// 读取客户端数据
	for {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		// fmt.Println("read:",read)
		if err != nil {
			fmt.Println("connect.Read err", err)
			return
		}

		// 用来查看自动发送的内容
		//fmt.Println("read content is" ,int(buf[0]))
		//fmt.Println("read content is" ,int(buf[1]))
		go func() {
			if int(buf[0]) != 255 {
				if int(buf[0]) != 13 {
					if string(buf[:6]) == "rename" {
						newUser.name = string(buf[7:read])
					} else {
						// message <- newUser.name
						messageToClient := fmt.Sprint(newUser.name + ":" + string(buf[:read]))
						fmt.Println("messageToClient:", messageToClient)
						message <- messageToClient
					}
				}
			}
		}()
		go func() {
			//speak_name := <- newUser.msg
			messageInfo := <-newUser.msg
			sp := fmt.Sprint(messageInfo + "\r\n")
			_, err2 := conn.Write([]byte(sp))
			if err2 != nil {
				fmt.Println("write data to client err", err2)
			}
			fmt.Println("client messageInfo", messageInfo)
		}()

	}
}

func main() {
	// 创建服务
	listen, err := net.Listen("tcp", ":8011")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	fmt.Println("server start success")

	//启动全局唯一go程,监听message通道,写给所有用户
	go broadcast()

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
		time.Sleep(500)
	}

}
