package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int -> float
	var x = 1
	xx := float64(1)
	fmt.Printf("%T %v %f\n", x, x, x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	// float -> int
	var y = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	// string -> int
	var s = "14"
	//z := int(s) // できない
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("%T %v %d\n", i, i, i)
}
