package main

import "fmt"

const (
	c1 = iota
	c2
	c3
)

const (
	_      = iota
	KB int = 1 << (10 * iota) // 1024
	MB                        // 1048576
	GB                        // 1073741824
)

func main() {
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)
}
