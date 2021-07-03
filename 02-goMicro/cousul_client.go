package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"socket/golearn/02-goMicro/pb"
)



func main() {

	conn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)

	person := pb.Person{
		Name: "Andy",
		Age:  20,
	}

	p, err := client.SayHello(context.TODO(), &person)
	fmt.Println(p)
}