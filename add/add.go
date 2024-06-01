package add

import "fmt"

func init() {
	fmt.Println("first init()")
}

func init() {
	fmt.Println("second init()")
}

func Add(a, b int) int {
	return a + b
}
