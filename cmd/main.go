package main

import (
	"fmt"
	// "./animals"
	"math"
)

func main() {
	// fmt.Println("Hello Go!")

	// fmt.Println(animals.ElephantFeed())
	// fmt.Println(animals.MonkeyFeed())
	// fmt.Println(animals.RabbitFeed())
	
	// fmt.Println(AppName())

	// a := [3] string {
	// 	"Taro",
	// 	"Hanako",
	// 	"Kenji",
	// }
	// fmt.Println(a)

	// int型の変数nを定義
	// varを使って変数を定義する場合、変数の名前と型の両方を明示的に指定する必要がある
	// var n int

	// まとめて定義も可能
	// var x, y, z int

	// 異なる型の変数をまとめて定義可能
	// var (
	// 	a, b int
	// 	c string
	// )

	// n = 5
	// n = "str"  // 異なる型の格納はコンパイルエラーを起こす

	// 複数の値をまとめて格納可能
	// x, y, z = 1, 2, 3
	// x, y = 4, 5, 6  // 左辺と右辺の個数に差があるとコンパイルエラーを起こす

	// :=を使って代入すると、型推論が起こる(暗黙的な定義)
	// i := 1

	// 暗黙的な変数定義は一度しか許されない
	// あくまで変数定義であって、型推論や初期値の代入は副次的なものだと理解しておく
	// e := 1
	// e := 2  // コンパイルエラー

	// varを利用した明示的な変数定義の場合でも、エラーが発生する
	// 明示的、暗黙的に関わらず「同じ変数を複数回定義するとエラーが発生する」という原則は押さえておく
	// var s string
	// var s string  // コンパイルエラー

	// 変数への再代入には制限がない
	// var r int
	// r = 1
	// r = 2

	// c := 1
	// c = 2

	// fmt.Println(increment())
	// fmt.Println(decrement())
	
	// fmt.Println(local_increment())
	// fmt.Println(local_decrement())

	// a := 1
	// b := 2
	// c := 3
	// fmt.Println(a)  // bとcが宣言されているものの使っていないので、エラーが出る

	// m, n := 1, 2
	// fmt.Printf("m = %d\nn = %d\n", m, n)

	// 整数の型変換(キャスト)
	// n := uint(17) // uint型で定義
	// b := byte(n)  // byte型へ変換
	// i64 := int64(n)  // int64型へ変換
	// u32 := uint32(n)  // uint32型へ変換

	// fmt.Printf("uint32 max value = %d\n", math.MaxUint32)
	// fmt.Printf("int64 max value = %d\n", math.MaxInt64)
	// fmt.Printf("int64 min value = %d\n", math.MinInt64)
}