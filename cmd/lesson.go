package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	/*
		// channelから1つとり出す
		x := <-ch
		fmt.Println(x)
		fmt.Println(len(ch))

		ch <- 300
		fmt.Println(len(ch))
	*/

	close(ch) // channelをcloseして終了させる

	for c := range ch {
		fmt.Println(c)
	}
}
