package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, i int) {
	// something
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	// producerでchに値が挿入される度に反応
	// ch がcloseするまで待機中のイメージ
	for i := range ch {
		// deferでwg.Done()を定義させるためにインナー関数にする
		// -> 処理の途中でエラーになった場合でもwg.Done()が実行されるように
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
	fmt.Println("################")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer（chに挿入）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer（chの中身を処理）
	go consumer(ch, &wg)
	wg.Wait()
	close(ch) // consumer関数のrange chを終了させる
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
