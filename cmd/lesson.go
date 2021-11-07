package main

import "fmt"

type Vertex struct {
	// 大文字はpublic
	X, Y int
	S    string
	// 小文字はprivateの意
	x int
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000 // structの場合*を付けなくても実態にアクセスされる
}

func main() {
	vv := Vertex{X: 1, Y: 2}
	changeVertex(vv)
	fmt.Println(vv) // 書き変わらない
	// {1 2  0}

	vv2 := &Vertex{X: 1, Y: 2}
	changeVertex2(vv2)
	fmt.Println(*vv2) // 書き換わる
	// {1000 2  0}

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test", 1}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	var v5 Vertex
	fmt.Println(v5) // スライスやマップはnilだがstructは0

	v6 := new(Vertex)
	fmt.Println(v6) // これはポインタが返る

	v7 := &Vertex{}
	fmt.Println(v7) // これもポインタが返る
}
