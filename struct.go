package main

import "fmt"

type student struct {
	name  string
	age   int
	score float64
}

func main() {

	//	fmt.Println("hello world")
	lily := student{
		name:  "lily",
		age:   20,
		score: 90,
	}

	Bob := student{
		"bob",
		21,
		99,
	}

	fmt.Println(lily)
	fmt.Println(Bob)

	ll := &lily

	fmt.Println(ll.name, (*ll).age)
}
