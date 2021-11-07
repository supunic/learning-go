package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Printf("%T %v\n", s, s)

	m := make(map[string]int)
	fmt.Printf("%T %v\n", m, m)

	var p = new(int) // メモリの領域を確保する
	fmt.Printf("%T %v\n", p, p)

	fmt.Println(p)  // 0xc000018100
	fmt.Println(*p) // 0
	*p++
	fmt.Println(*p) // 1

	ch := make(chan int)
	fmt.Printf("%T\n%v\n", ch, ch)

	st := new(struct{})
	fmt.Printf("%T\n%v\n", st, st)

	// makeもnewもメモリの領域を確保する
	// make -> 値を返す
	// new -> ポインタを返す

	/*
		var p2 *int
		fmt.Println(p2) // <nil>
	*/
}
