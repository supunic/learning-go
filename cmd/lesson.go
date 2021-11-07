package main

import "fmt"

func thirdPartyConnectDB() {
	// golangはpanic非推奨
	// 自分のコードでpanicを起きないようにかけ
	panic("Unable to connect database!")
}

func save() {
	// thirdPartyConnectDB()
	defer func() {
		s := recover() // panicをrecoverがキャッチする
		fmt.Println(s) // 例外メッセージの出力
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("ok")
}
