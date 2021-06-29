package main

import (
	"fmt"
	"net"
)

/*
1.有很多用户
2. 有个user结构，主要包含msg管道
3. 有个进行全局广播的管道
4. 单独go程服务器获取用户内容后 发送到全局用户中
5.全局的map管理所有的user
*/

type User struct {
	id string
	name string
	msg chan string

}
var AllUser = make(map[string]User)

var message = make(chan string)

func broadcase()  {
	/*
	向全部用户广播, 全局唯一go程
	 */
	info := <-message
	fmt.Println("broadCast go rounte start suss")
	for _,user := range AllUser{
		user.msg <- info
	}
}


func handle(conn net.Conn)  {
	fmt.Println("start handle success")
	ClientInfo := conn.RemoteAddr().String()
	newUser := User{
		id:   ClientInfo,
		name: ClientInfo,
		msg:  make(chan string),
	}
	go func() {
		messageInfo := <- newUser.msg
		fmt.Println(messageInfo)
	}()
	AllUser[newUser.id] = newUser
	sprintf := fmt.Sprintf("%s:%s ====> ONLINE", newUser.name, newUser.id)
	message <- sprintf
	for{
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Println("connect.Read err",err)
			return
		}
		fmt.Println("read content is " ,string(buf[:read]))

	}
}

func main() {
	// 创建服务
	listen, err := net.Listen("tcp", ":8011")
	if err != nil {
		fmt.Println("listen err:",err)
		return
	}
	fmt.Println("server start success")

	//启动全局唯一go程,监听message通道,写给所有用户
	go broadcase()

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
