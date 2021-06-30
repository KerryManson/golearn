package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8023")
	if err != nil {
		fmt.Println("Dial Failed :", err)
		return
	}
	defer client.Close()
	var reply string
	err = client.Call("hello.HelloWorld", "大哥", &reply)
	if err != nil {
		fmt.Println("Call Failed", err)
	}
	fmt.Println(reply)
}
