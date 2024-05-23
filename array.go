package main

import "fmt"

func main() {
	nums := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	//fmt.Println()
	for key, value := range nums {
		value++
		fmt.Println(key, value)
	}
	names := []string{"a", "b", "c", "d", "e", "f"}

	for i, v := range names {
		names = append(names, v)
		fmt.Println("index:", i, "value:", v, "capacity:", cap(nums))
	}

	n := []int{}
	for i := 0; i < 30; i++ {
		n = append(n, i)
		fmt.Println("index:", i, "capacity:", cap(n))
	}

	fmt.Println("---------------------")

	n2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n3 := n2[:]
	//n3[0] = 10
	fmt.Println(n3)
	//fmt.Println(n2)

	str := "helloworld"[5:7]
	fmt.Println(str)

	fmt.Println("---------------------")

	n4 := []int{1, 2, 3}
	n5 := make([]int, 3, 20)
	copy(n4, n2[:])
	copy(n5, n2[:])
	fmt.Println(n4)
	fmt.Println(n5)
}
