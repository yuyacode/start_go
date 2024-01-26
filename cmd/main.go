package main

import (
	"fmt"
	// "./animals"
	// "math"
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

	// float32型の最大と最小、float64型の最大と最小を出力
	// fmt.Println(math.MaxFloat32)
	// fmt.Println(math.SmallestNonzeroFloat32)
	// fmt.Println(math.MaxFloat64)
	// fmt.Println(math.SmallestNonzeroFloat64)

	// f64 := 1.0  // 32ビットアーキテクチャだったとしても、float64型になる (intとは異なる)
	// f32 := float32(1.0)  // float32型 (float32型で定義したい場合は、float32()を使った明示的な型変換が必要)

	// zero := 0.0
	// pinf := 1.0 / zero   // +Inf 正の無限大
	// ninf := -1.0 / zero  // -Inf 負の無限大
	// nan := zero / zero   // NaN　非数

	// fmt.Println("\u65e5 本 \U00008a9e")  // 日本語と出力
	// fmt.Println("\xff\u00FF")  // 特殊文字を出力

	// ``を使った複数行に渡る文字列を「RAW文字列リテラル」と表現されている
	// s := `
	// Goの
	// RAW文字列リテラルによる
	// 複数行に渡る
	// 文字列
	// `
	// fmt.Printf("%v", s)

	// a := `
	// \n
	// \n
	// `
	// fmt.Printf("%v", a)  // RAW文字列リテラルでは、\nは改行ではなく、そのまま文字として扱われる。名前の通り、RAW(生)である

	// a := [5] int {1, 2, 3, 4, 5}
	
	// fmt.Printf("%d\n", a[0])  // 1
	// fmt.Printf("%d\n", a[1])  // 2
	// fmt.Printf("%d\n", a[2])  // 3
	// fmt.Printf("%d\n", a[3])  // 4
	// fmt.Printf("%d\n", a[4])  // 5

	// b := [5] int {}         // 要素が少ない分には問題ない {0, 0, 0, 0, 0} となる
	// c := [5] int {1, 2, 3}  // 要素が少ない分には問題ない {1, 2, 3, 0, 0} となる
	// d := [5] int {1, 2, 3, 4, 5, 6}  // 多いとエラーとなる

	// intの配列において、初期値を与えなかったり、空配列を初期値とした場合、要素には0が割り当てられる
	// var a [3] int    // {0, 0, 0}
	// a := [3] int {}  // {0, 0, 0}

	// boolの場合は、[false, false, false] となる
	// var a [3] bool
	// a := [3] bool {}

	// stringの場合、["", "", ""] となる
	// var a [3] string
	// a := [3] string {}

	// 要素数の省略が可能
	// array_1 := [...] int {1, 2, 3}  // [3]int型
	// array_2 := [...] int {1, 2, 3, 4, 5}  // [5]int型
	// array_3 := [...] int {}  // [0]int型

	// 要素代入時、型と互換性がない場合はエラー
	// arr := [...] uint8 {1, 2, 3}
	// arr[0] = 256  // uint8は、0〜255まで扱い可能なので (constant 256 overflows uint8)
	
	// var (
	// 	arr1 [3] int
	// 	arr2 [5] int
	// )
	// arr1 = arr2  // 要素数が異なるのでエラーになる。要素数が異なるだけで全く別の型として扱われる

	// var (
	// 	arr_1 [4] int
	// 	arr_2 [4] uint
	// )
	// arr_1 = arr_2  // 当然、要素数が同じだとしても、型が違えば、全く別の型として扱われる

	// var x interface{}
	// fmt.Printf("%#v\n", x)  // <nil>  初期化されていない場合は<nil>が入る  <nil>とは「具体的な値を持っていない状態」を表す特殊な値  他言語におけるnullのようなもの

	// var (
	// 	a, b, c, d, e interface{}
	// )

	// a = 1
	// b = 3.14
	// c = '山'  // rune
	// d = "文字列"
	// e = [...] uint8 {1, 2, 3, 4}

	// interface{}型には、任意の型の値を格納可能。そのため実際の型が実行時まで不明である可能性が高いため、%vフォーマット指定子を使用して、異なる型に応じて適切な表示を行うことが一般的
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", d)
	// fmt.Printf("%v\n", e)

	// var (
	// 	num1, num2 interface{}
	// )
	// num1 = 100
	// num2 = 100
	// sun_num := num1 + num2  // interface{}型はあくまで全ての型の値を汎用的に表す手段であり、一旦interface{}型に値が格納されてしまうと、演算対象としては利用できないため注意。型情報を失ってしまうみたい。

	// +算術演算子は、文字列の結合にも使える
	// s := "Taro" + " " + "Tanaka"
	// fmt.Println(s)

	// 下記は、整数のビット演算を行う演算子
	// 論理積（AND） 整数同士の各ビットを比較して、双方が1である場合に1に、それ以外は0になるビット演算を行う
	// n := 165 & 155  // 129

	// 論理和（OR） 整数同士の各ビットを比較して、どちらかあるいは両方が1である場合に1に、双方が0である場合のみ0になるビット演算を行う
	// n := 197 | 169  // 237
	
	// 排他的論理和（XOR） 整数同士の各ビットを比較して、どちらか片方のみ1である場合に1に、それ以外は0になるビット演算を行う
	// n := 92 ^ 137  // 213

	// ビットクリア（AND NOT） X &^ Yは、X and (NOT Y)という２つのビット演算をまとめた短縮形
	// n := 108 &^ 13 // 96
	// n := 1 &^ 1 // 0  同じ整数間でビット演算を行うと0になる

	// 左シフト  整数のビット値を左へ指定した整数分だけずらすビット演算  0000 0001 が 0000 0010 になる  よって1から2になる
	// n := 1 << 1  // 2
	// n := 1 << 7  // 2

	// 右シフト  左シフトの逆  右へずらす
	// n := 218 >> 4  // 13

	// 下記のように省略可能
	// n += 5
	// s += " ありがとうございます "
	// n *= 10
	// n &= x
	// n <<= 3

   // fmt.Println(^int8(16))  // -17
   // fmt.Println(^uint8(16))  // 239

   // fmt.Println(plus(51, 37))

	// hello()

	// q, r := div(10, 3)
	// fmt.Printf("%d余り%d\n", q, r)

	// _を使って戻り値の一部を破棄することも可能
	// q, _ := div(10, 3)
	// fmt.Printf("%d\n", q)
	// _, r := div(10, 3)
	// fmt.Printf("%d\n", r)

	// 暗黙的に定義する変数が存在せずエラーになる
	// _, _ := div(10, 3)

	// 下記のような書き方なら可能だが、これはただ関数を呼ぶだけと同じなので、ただ冗長なだけ
	// _, _ = div(10, 3)
	// div(10, 3)

	// Goでは例外の機構がない。その代わり、戻り値の一部でエラー発生有無を返すことで、関数呼び出しが成功したかどうかを確認する
	// result, err := doSomething()
	// if (err != nil) {
	// 	// エラー処理
	// }

	fmt.Println(doSomething())  // 0
}

func doSomething() (a int) {
	return
}

// 上記は下記の省略形

func doSomething() int {
	var a int
	return a
}