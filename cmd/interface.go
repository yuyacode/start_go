package main

import (
	"fmt"
)

type myError struct {
	message string
	code int
}

// インターフェースは型の一種であり、任意の型が「どのようなメソッドを実装するべきか」を規定するための枠組み
// interface{メソッドのシグネチャの列挙}という形式
// 型が実装するべきメソッドのシグネチャを任意の数だけ列挙可能

// Goの組み込み型 error
// error型（errorインターフェース）では、文字列とint型を返すメソッドErrorのみ定義されている
// これは型エイリアスではなく、errorという名前のインターフェース型を定義している
type error interface{
	Error() (string, int)
}

// errorインターフェースが要求するErrorメソッドを定義
func (e *myError) Error() (string, int) {
	return e.message, e.code
}

func interfaceFunc() {
	err := RaiseError()  // errは、実際には*myError型だが、プログラム上ではあくまでerror型
	errorMessage, errorCode := err.Error()
	fmt.Println(errorMessage)
	fmt.Println(errorCode)

	// もし変数errから本来の型である*myError型を取り出したい場合は、型アサーションを使用する
	e, ok := err.(*myError)
	if ok {
		e.message  // エラーが発生しました
		e.code     // 500
	}
}

func RaiseError() error {
	return &myError{message: "エラーが発生しました", code: 500}
}