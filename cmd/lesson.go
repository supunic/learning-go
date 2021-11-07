package main

import (
	"fmt"
	"time"
)

func main() {
	// 100 milli secごとにchanを返す
	tick := time.Tick(100 * time.Millisecond)
	// 500 milli sec後にchanを返す
	boom := time.After(500 * time.Millisecond)

OuterLoop:
	for {
		select {
		case <-tick: // time.Tickの戻り値chanを受信
			fmt.Println("tick.")
		case <-boom: // time.Afterの戻り値chanを受信
			fmt.Println("BOOM!")
			break OuterLoop
		default:
			// tickもboomもchに入ってなかったときにforが回ってきた場合
			// 50 milli sec 待ってまたforが始まる
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("OuterLoop broken")
}
