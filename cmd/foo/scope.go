package foo

import (
//
)

const (
	Max            = 100
	internal_const = 1
)

func FooFunc(n int) int {
	return internalFunc(n)
}

func internalFunc(n int) int {
	return n + 1
}

// 日本語など大文字、小文字のない文字の場合、小文字と同様に他パッケージから参照できない

func sum(a, b int) (result int) {
	// var result int や result := a + b のように再度宣言する必要はない
	// 'result' はすでに定義されているので、直接値を代入できる
	result = a + b
	return // 'return' のみで 'result' の現在の値が返される
}

// func bar(a int) (b string) {
// 	var a int           // aは定義済みなのでエラーになる
// 	const b = "string"  // bは定義済みなのでエラーになる
// 	return b
// }

func Hoge() (b string) { // そもそも識別子の重複は避けるべき
	b = "string"
	{
		// 関数より深いブロック
		const b = "str" // ブロック内でのみ有効な定数
		// 中略
	}
	return b // string
}

func Biz() (b string) {
	b = "string"
	{
		// 関数より深いブロック
		const b = "str"
		return // return文は関数レベルでのみ有効であり、ブロックレベルでは許可されていない。コンパイルエラー
		// 中略
	}
	return // 外に出す
}
