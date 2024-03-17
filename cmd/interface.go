package main

import (
	"fmt"
)

type myError struct {
	message string
	code int
}
type person struct {
	name string
	age int
}
type car struct {
	number string
	model string
}

// インターフェースは型の一種であり、任意の型が「どのようなメソッドを実装するべきか」を規定するための枠組み
// interface{メソッドのシグネチャの列挙}という形式
// 型が実装するべきメソッドのシグネチャを任意の数だけ列挙可能
// インターフェースとは、異なる型をひとつにまとめたり、グルーピングする役割を持つ抽象的な型
// インターフェースを利用することで、異なる具体的な型が共通の振る舞いを持つことができる
// 具体的な型は、インターフェースに定義された全てのメソッドを実装する必要がある。全て実装することで、そのインターフェースを実装したと見なされる

// Goの組み込み型 error
// error型（errorインターフェース）では、文字列とint型を返すメソッドErrorのみ定義されている
// これは型エイリアスではなく、errorという名前のインターフェース型を定義している
type error interface {
	Error() (string, int)
}
type stringify interface {
	toString()
}

// errorインターフェースが要求するErrorメソッドを定義
func (e *myError) Error() (string, int) {
	return e.message, e.code
}
// stringifyインターフェースが要求するtoStringメソッドを定義
// person型とcar型という２つの全く異なる型それぞれに、stringifyインターフェースで定義されたtoStringメソッドを実装
func (p *person) toString() {
	fmt.Println(fmt.Sprintf("%s(%d)", p.name, p.age))  // taro(21)
}
func (c *car) toString() {
	fmt.Println(fmt.Sprintf("[%s] %s", c.number, c.model))  // [ABC-1234] GB66y
}

func interfaceFunc() {
	err := RaiseError()  // errは、実際には*myError型だが、プログラム上ではあくまでerror型
	errorMessage, errorCode := err.Error()
	fmt.Println(errorMessage)
	fmt.Println(errorCode)

	// もし変数errから本来の型である*myError型を取り出したい場合は、型アサーションを使用する
	// e, ok := err.(*myError)
	// if ok {
	// 	e.message  // エラーが発生しました
	// 	e.code     // 500
	// }

	vs := []stringify {
		&person{name: "taro", age: 21},
		&car{number: "ABC-1234", model: "GB66y"},
	}
	for _, v := range vs {
		v.toString()
	}

	Println(&person{name: "田中", age: 23})           // 田中(23)
	Println(&car{number: "U1234", model: "j-3456"})  // [U1234] j-3456
}

func RaiseError() error {
	return &myError{message: "エラーが発生しました", code: 500}
}

// インターフェースを活用すれば、汎用性の高い関数やメソッドを定義可能
// 下記のPrintln関数は、stringifyインターフェースさえ実装していれば、どのような型からでも呼び出すことができる
// このように「型の性質」を抽出したインターフェースを定義すれば、Goの厳密な型システムに緩やかな柔軟性を与えることが可能になる
func Println(s stringify) {
	s.toString()
}