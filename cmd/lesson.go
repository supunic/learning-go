package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	// Vertex ... メソッドの値レシーバ
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	// *Vertex ... メソッドのポインタレシーバ
	v.x = v.x * i
	v.y = v.y * i
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}
