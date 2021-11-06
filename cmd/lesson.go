package main

import "fmt"

var (
	i    = 1
	f64  = 1.2
	s    = "test"
	t, f = true, false
)

func foo() {
	// short valuable declarations
	// -> 関数内のみ
	xi := 1                // := で宣言と代入、型を自動で割り振り
	xf64 := 1.2            // デフォルトでfloat64になる
	var xf32 float32 = 1.2 // float32に指定した場合
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xf32)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
