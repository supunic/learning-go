package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	// Vertex ... メソッドの値レシーバ
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	// *Vertex ... メソッドのポインタレシーバ
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex) int {
	// ただの関数
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	//fmt.Println(Area(v))
	v.Scale(10)
	fmt.Println(v.Area()) // method
}
