package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	/*
		// 無限ループ
		for {
			fmt.Println("hello")
		}
	*/
}
