package main

import (
	"context"
	"google.golang.org/grpc"
	pb "socket/golearn/01-microServer/pd"
)
// 定义类
type Children struct {

}

func (this *Children) SayHello(context.Context, *Teacher) (*Teacher, error) {

}

func main() {
	// 初始化一个grpc对象
	grpcServer := grpc.NewServer()

	// 2.注册服务
	pb.RegisterSayNameServer()




}