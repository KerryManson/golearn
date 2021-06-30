package main

import (
	"net/rpc"
)

// 目的 要求服务器在注册rpc对象时,编译器就能检测出注册对象是否合法
// 多态的体现

//  创建接口,在接口中定义方法原型
//  服务端
type MyInterface interface {
	HelloWorld(string, *string) error
}

// 调用该方法时, 需要给 I传参, 参数应该是,实现了 helloworld方法的类对象
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

// 客户端
type MyClient struct {
	c *rpc.Client
}

func ClientInit(addr string) MyClient {
	client, _ := rpc.Dial("tcp", addr)
	return MyClient{c: client}
}

func (this MyClient) HelloWorld(a string, b *string) error {
	err := this.c.Call("hello.HelloWorld", a, b)
	return err

}
