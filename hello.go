package main

import "fmt"

func main() {
	i := 1
	i++
	fmt.Println(i)
	fmt.Println("hello world")
	_, _ = test(1, "aaa")

	fmt.Println("---------------------------")

	ptr := &i
	fmt.Println(ptr)
	fmt.Println(*ptr)

	ptr2 := new(string)
	*ptr2 = "123"
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	fmt.Println("------------------------")

	res := test2()
	fmt.Println(*res)

	if res == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	const ( // 常量组
		size int64 = 1024
		eof        = -1 // 整型常量, 自动推导类型
	)

	var (
		v5 int
		v6 int
	)
	fmt.Println(size, eof)
	fmt.Println(v5, v6)
}

func test(a int, b string) (int, string) {
	fmt.Println(a)
	fmt.Println(b)
	return a, b
}

func test2() *string {
	city := "xian"
	ptr := &city
	return ptr
}
