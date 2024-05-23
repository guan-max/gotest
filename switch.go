package main

import "fmt"

func main() {

	//cmd := os.Args
	//
	//for i, v := range cmd {
	//	fmt.Println(i, v)
	//}
	//if len(cmd) < 2 {
	//	fmt.Println("args error")
	//	return
	//}
	//
	//switch cmd[1] {
	//case "1":
	//	fmt.Println("1")
	//	fallthrough
	//case "2":
	//	fmt.Println("2")
	//case "3":
	//	fmt.Println("3")
	//default:
	//	fmt.Println("default")
	//}
	//fmt.Println("Hello World")
	//i := 3

	//LABEL121:
	//	fmt.Println("label121:")
	//	if i == 3 {
	//		i++
	//		goto LABEL121
	//	}
	//for i := 0; i < 5; i++ {
	//	for j := 0; j < 5; j++ {
	//		if j == 3 {
	//			//goto LABEL121
	//			//continue LABEL121
	//			break LABEL121
	//			//break
	//		}
	//
	//		fmt.Println("i:", i, ",j:", j)
	//	}
	//}
	const (
		MONDAY = iota + 1
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
	)

	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDAY)
	//fmt.Println("over!")
}
