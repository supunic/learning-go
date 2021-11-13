package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()         // メモリの参照するアドレス先をロック
	defer c.mux.Unlock() // 終わったらロック解除
	c.v[key]++           // 値を書き換える
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)}

	// 2つのgoroutineがメモリ参照で競合しない
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()

	time.Sleep(1 * time.Second)

	fmt.Println(c, c.Value("key"))
}
