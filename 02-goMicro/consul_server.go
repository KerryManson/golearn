package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
	"socket/golearn/02-goMicro/pb"
)

type Worker struct {

}

func (worker *Worker) SayHello(ctx context.Context, p *pb.Person)  (*pb.Person,error){
	p.Age +=10
	fmt.Println("age is :",p.Age)
	return p,nil
}

func main()  {

	// 把grpc服务注册到consul上
	// 1. 初始化consul配置
	defaultConfig := api.DefaultConfig()

	// 2. 创建consul对象
	newClient, err2 := api.NewClient(defaultConfig)
	if err2 != nil {
		fmt.Println("consul NewClient err:",err2)
		return
	}
	// 3.设置监听 制定IP 端口
	serviceRegistration := api.AgentServiceRegistration{
		Kind:    "",
		ID:      "jw",
		Name:    "HelloService",
		Tags:    []string{"TestHello"},
		Port:    8800,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			TCP:      "127.0.0.1:8800",
			Timeout:  "5s",
			Interval: "5s",
		},
	}

	// 4. 注册grpc到consul上
	newClient.Agent().ServiceRegister(&serviceRegistration)



	grpcserver := grpc.NewServer()

	pb.RegisterHelloServer(grpcserver,new(Worker))

	listen, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		return
	}
	defer listen.Close()

	fmt.Println("开启服务")
	grpcserver.Serve(listen)
}
