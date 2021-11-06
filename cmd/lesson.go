package main

import "fmt"

// mainより先に呼び出される
func init() {
	fmt.Println("Init!")
}

func buzz() {
	fmt.Println("Buzz")
}

func main() {
	buzz()
	fmt.Println("Hello world!", "test test")
}
