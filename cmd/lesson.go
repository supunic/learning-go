package main

import "fmt"

func one(x *int) {
	// ポインタxの中身に1を入れている
	*x = 1
}

func main() {
	// ポインタ

	var n = 100
	one(&n)
	fmt.Println(n) // 値の表示

	fmt.Println(&n) // 値を入れたメモリのアドレスを表示

	var p = &n // *int intのポインタ型
	fmt.Println(p)
	fmt.Println(*p) // ポインタの中身を表示
}
