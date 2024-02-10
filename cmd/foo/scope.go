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
	// resultは戻り値として既に定義されているので、直接値を代入できる
	result = a + b
	return // returnのみでresultの現在の値が返される
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
		// 追記: そんなことない。returnはブロックレベルでも使用可能。ブロック内でのreturnは関数レベルでのreturnと同様の役割をする。すなわち関数から抜けつつ、戻り値を返す
		// 関数ブロックと同様の役割を果たすため、strが格納された定数bではなく、stringが格納された変数bを返す
	}
	// return // 外に出す  不要
}
