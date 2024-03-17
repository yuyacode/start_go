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

type t struct {
	id int
	language string
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

// インターフェースのメソッドに関しても、大文字から始まる場合は外部パッケージからアクセス可能
// インターフェースのメソッドを非公開にし、外部から隠蔽することによるメリットはほとんどない。理由は下記。
// ・インターフェースの主な目的の一つは、異なる型が共通の振る舞いを持つことを抽象化し、定義すること。これにより、異なる型を同一のインターフェースとして扱うことが可能になり、多様性（ポリモーフィズム）を提供する。
//   インターフェースのメソッドを非公開にすると、そのインターフェースをパッケージ外で実装する能力が制限され、インターフェースの利用価値が大幅に低下する

// インターフェースの内部にインターフェースを含める定義も可能
// I1が要求するメソッドは、Method0とMethod1の２つ
// I2が要求するメソッドは、Method0とMethod1とMethod2の３つ
// メソッド名の重複には注意
type I0 interface {
	Method0() int
}
type I1 interface {
	I0
	Method1() int
}
type I2 interface {
	I1
	Method2() int
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

	t := &t{id: 1, language: "日本語"}
	fmt.Println(t)  // <<1, 日本語>>
	// このコードが正常に動作する理由は、fmtパッケージのPrintln関数（および他の出力関数）が、出力する値（今回であれば変数t）に対して、fmt.Stringerインターフェースが実装されているかどうかを内部的にチェックしているから
	// もし実装していれば、そのStringメソッドを呼び出し、返された文字列を出力する。この振る舞いにより、カスタム型が自身をどのように文字列化して表示するかを制御することができる
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

// fmtパッケージに定義されているStringerインターフェース
// type Stringer interface {
//     String() string
// }

func (t *t) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.id, t.language)
}