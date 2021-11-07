package main

import "fmt"

func do(i interface{}) {
	// intやstringをinterfaceで受け取る

	//ii := i.(int) // タイプアサーション
	//ii *= 2
	//fmt.Println(ii)

	//ss := i.(string) // タイプアサーション
	//fmt.Println(ss)

	// switch type 文
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	do(10)
	do("Mike")
	do(true)

	var i = 10
	ii := float64(10)
	fmt.Println(i, ii)
}
