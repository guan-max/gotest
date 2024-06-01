package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义类对象
type World struct {
}

// 绑定类方法

func (this *World) HelloWorld(name string, rsp *string) error {
	*rsp = "Hello " + name
	return nil
}

func main() {
	//1 注册rpc服务，绑定类方法
	rpc.RegisterName("hello", new(World))

	//2 设置监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port 1234")
	//3 建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Println("Accepted connection")
	//4 绑定服务
	rpc.ServeConn(conn)
}
