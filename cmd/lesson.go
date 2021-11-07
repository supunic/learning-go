package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
}

func main() {
	foo()
	foo(10, 20)
	foo(10, 20, 30)
	// 0 []
	// 2 [10 20]
	// 3 [10 20 30]

	s := []int{1, 2, 3}
	fmt.Println(s)
	// [1 2 3]

	foo(s...)
	// 3 [1 2 3]
}
