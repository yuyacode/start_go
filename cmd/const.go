package main

import (
	"fmt"
	"math"
)

const ONE = 1

// 変数同様、まとめて定義可能
// const (
// 	X = 1
// 	Y = 2
// 	Z = 3
// )

const (
	// 定数の値を省略することができる
	X = 1  // X = 1
	Y      // Y = 1
	Z      // Z = 1
	S1 = "あ"  // S1 = "あ"
	S2         // S2 = "あ"

	// 上記のような「型なし定数」は、特定の型を持たない定数。コンパイル時に適切な型が推論されるため柔軟。適切な型を取ろうとする関係上、同じ定数が異なる状況で異なる型を持つことがある
	// const X = 1  これがint64として扱われるケースもあれば、uint64として扱われるケースもある。その状況により適切な型が取られる
	// その定数が代入される変数の型や、関数の引数の型、演算のコンテキストなどに基づいて、適切な型に「推論」される

	// 型あり定数の定義方法
	I64 int64 = -1  // I64はint64型
	F64 float64 = 1.2  // F64はFloat64型

	// 下記のような記述も可能
	// I64 = int64(-1)
	// F64 = float64(1.2)

	uint64_12345 = uint64(12345)
)

func main() {
	one, two := one()
	fmt.Printf("%d, %d\n", one, two)

	// var int64_12345 int64
	// int64_12345 = uint64_12345  // 型が異なるため、コンパイルエラー

	i := uint64_12345  // uint64型だと推論される。代入先の変数の型は、定数と同じ型であると推論される
	fmt.Println(i)       // 12345
	fmt.Printf("%T\n", i)  // uint64

	f32 := float32(math.Pi)
	f64 := float64(math.Pi)
	fmt.Printf("%v\n", f32)
	fmt.Printf("%v\n", f64)

	// float64型などの基本型を介さず、定数間で演算を行うことで、浮動小数点の丸めによる誤差を抑えることが可能
	const F = 1.0000000000001
	fmt.Println(float64(F) * 10000)  // 10000.000000000999
	fmt.Println(F * 10000)  // 10000.000000001
}

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

