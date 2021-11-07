package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done() // goroutineの終了の宣言
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	// sync.WaitGroup ... goroutineの制御に使用する
	var wg sync.WaitGroup
	defer wg.Wait() // goroutine待機の宣言

	wg.Add(1)                  // goroutine追加
	go goroutine("world", &wg) // goroutine実行

	normal("hello")
}
