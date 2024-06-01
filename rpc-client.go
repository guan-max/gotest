package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1 rpc 连接服务器
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	//2 调用函数
	//传出参数
	var res string
	err = client.Call("hello.HelloWorld", "123", &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
