package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b = [2]int{100, 200}
	// 配列はサイズ変更できない
	//b = append(b, 300)
	fmt.Println(b)
}
