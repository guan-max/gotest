package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("map.go")

	//匿名函数
	//defer调用有返回值的函数，使用匿名函数
	defer func() {
		fmt.Println("defer")
		_ = file.Close()
	}()
	fmt.Println(err)
	buf := make([]byte, 1024)

	defer fmt.Println("0000000")
	defer fmt.Println("1111111")
	defer fmt.Println("2222222")

	n, e := file.Read(buf)
	fmt.Println(n, e)
	fmt.Println(string(buf[:n]))
}
