package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// -> len=3 cap=5 value=[0 0 0]

	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// -> len=5 cap=5 value=[0 0 0 0 0]

	// capを超えても追加できる
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// -> len=10 cap=10 value=[0 0 0 0 0 1 2 3 4 5]

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	// -> len=3 cap=3 value=[0 0 0]

	b := make([]int, 0, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	// -> len=0 cap=5 value=[]

	c := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
	}
	// -> len=1 cap=5 value=[0]
	//    len=2 cap=5 value=[0 1]
	//    len=3 cap=5 value=[0 1 2]
	//    len=4 cap=5 value=[0 1 2 3]
	//    len=5 cap=5 value=[0 1 2 3 4]

	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
	// -> len=5 cap=5 value=[0 1 2 3 4]
}
