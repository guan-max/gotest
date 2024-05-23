package main

import "fmt"

func main() {
	var stu map[string]int
	stu = make(map[string]int, 10)
	student := make(map[string]int, 10)
	fmt.Println(student)
	fmt.Println(stu)

	student["guangzhou"] = 100
	student["shanghai"] = 200

	fmt.Println(student)

	for k, v := range student {
		fmt.Println(k, v)
	}

	id := student["zzz"]
	fmt.Println(id)

	id, ok := student["zaa"]
	fmt.Println(id, ok)

	delete(student, "guangzhou")
	fmt.Println(student)

	b := test1(1, 2, "a")
	fmt.Println(b)

	p := testPtr1()
	fmt.Println(*p)
}

func test1(a, b int, c string) int {
	//return a, true
	//d = a
	//e = true
	return 0
}

func testPtr1() *string {
	name := "Duke"
	p0 := &name
	fmt.Println("*p0:", *p0)

	city := "深圳"
	ptr := &city
	return ptr
}
