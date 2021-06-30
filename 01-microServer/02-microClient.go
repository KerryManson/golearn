package main

import "fmt"

func main() {
	myClient := ClientInit("127.0.0.1:8023")
	var resp string
	err := myClient.HelloWorld("Liky", &resp)
	if err != nil {
		fmt.Println("Client Call Err:", err)
		return
	}
	fmt.Println(resp)

}
