package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println("Hello World"[0])         // 72 ascii code
	fmt.Println(string("Hello World"[0])) // H

	var s = "Hello World"
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(`
		Test
		Test
		Test
	`)
	fmt.Println("\"")
	fmt.Println(`"`)
}
