package main

import "fmt"

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	// eはポインタで受け取るのが良しとされている
	// -> エラー比較時に区別させるため
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false // something wrong ...
	if ok {
		return nil // 本来だとerrorなし
	}
	// &UserNotFound &つける
	return &UserNotFound{Username: "mike"}
}

func main() {
	// カスタムエラー
	// -> 自分で定義したエラーを出せる
	//    structを自分で用意して、そのstructにError()メソッドを追加する
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
