package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
	// 遅延実行 -> main関数が終了した後に実行される
	/*
		defer fmt.Println("world")
		foo()
		fmt.Println("hello")
		// hello
		// world
	*/

	/*
		// stacking defer
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
		// run
		// success
		// 3
		// 2
		// 1
	*/

	file, _ := os.Open("./lesson.go")
	defer file.Close()        // deferで忘れずにfileをclose
	data := make([]byte, 100) // バイト配列
	file.Read(data)           // バイト配列にデータを入れる
	fmt.Println(string(data)) // stringにcast
}
