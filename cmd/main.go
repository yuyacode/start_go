package main

import (  // import定義は、ファイル毎に独立しているため、同じmainパッケージに属するapp.goで定義しているfmtによるインポートは参照できない
	f "fmt"  // fで上書き可能、エイリアスではなく上書きなので、fmtは使用できなくなる
	// "./animals"
	// . "math"  // パッケージ名の省略が可能
	// "./foo"
)

func main() {

	// switchのcaseに式を与えることも可能
	n := 4
	switch {
	case n > 0 && n < 3 :
		f.Println("0 < n < 3")
	case n > 3 && n < 6 :
		f.Println("3 < n < 6")
	}

	// 定数を列挙したcaseと、式によるcaseの混在はエラーになる
	// 正しくは、定数を列挙したcaseはint型、式によるcaseはbool型となり、結果「型の不一致」が理由でコンパイルエラーが発生する
	switch x := 1; x {
	case 1, 2, 3 :
		f.Println(x)
	case x > 3 :
		f.Println("3より大きいです")
	}

	// 簡易文を使うことで、nはswitch文内でのみ有効な変数になる
	// 簡易文は、変数の局所性を高める重要な機能
	// switch n := 2; n {
	// case 1, 3, 5, 7, 9 :
	// 	f.Println("奇数")
	// case 2, 4, 6, 8, 10 :
	// 	f.Println("偶数")
	// }

	// n := 100  // 上記のswitchで使用されているnは、switch文内がスコープにため、ここで変数nを定義可能
	// f.Println(n)

	// n := 1
	// switch n {
	// case 1 :
	// 	f.Println("one")
	// case "2" :  // 型の互換性がないため、エラーになる
	// 	f.Println("two")
	// case 3 :
	// 	f.Println("three")
	// }

	// 型なし定数を使用してcase節を書くと、switch文に与えられた式の型と、case節の定数の型に互換性があるかチェックする
	// 型なし定数とは、明示的な型を持たない定数のこと。Goコンパイラによって特定の型に束縛されない
	// 型なし定数は、プログラムがコンパイルされる際に、使用される文脈に基づいて必要な型に自動的に適応する
	// num := 3
	// switch num {
	// case 1 :
	// 	f.Println("one")
	// case 2.0 :  // 浮動小数定数だが、整数2と互換性がある
	// 	f.Println("two")
	// case 3+0i :  // 複素数定数だが、整数3と互換性がある
	// 	f.Println("three")
	// }

	// 式を評価して分岐を行うswitch
	// n := 5
	// switch n {
	// 	case 1, 2 :
	// 		f.Println("1 or 2")
	// 	case 3, 4 :
	// 		f.Println("3 or 4")
	// 	default :
	// 		f.Println("unknown")
	// }

	// Goではbreakを書かなくとも、どこかのcaseでヒットすると次のcaseの評価へと移らない。すなわち、フォールスルーしない
	// フォールスルーしたい場合は、下記のように書く
	// fallthroughが定義されていると、次のcase内の処理を強制実行する。この際、次のcaseの評価は実行されない。すなわち、成立か不成立かの判定すらされず処理が実行される
	// s := "A"
	// switch s {
	// 	case "A" :
	// 		s += "B"  // "AB"
	// 		fallthrough
	// 	case "B" :
	// 		s += "C"  // "ABC"
	// 		fallthrough
	// 	case "C" :
	// 		s += "D"  // "ABCD"
	// 		fallthrough
	// 	default :
	// 		s += "E"  // "ABCDE"
	// }
	// f.Println(s)  // "ABCDE"

	// 範囲節によるfor
	// 範囲式は、予約語rangeと任意の式を組み合わせて定義する
	// fruits := [3]string {"Apple", "Banana", "Cherry"}
	// for i, s := range fruits {
	// 	f.Printf("fruits[%d] = %s\n", i, s)
	// }

	// 文字列型とrange
	// 要素には、Unicodeにおける文字コードが出力される
	// str := "abc"
	// for index, rune := range str {
	// 	f.Println(index)  // 0, 1, 2
	// 	f.Println(rune)  // 97, 98, 99
	// }

	// 文字に応じて、インデックスの増分量は異なる
	// str_ja := "あいう"
	// for index, rune_ja := range str_ja {
	// 	f.Println(index)  // 0, 3, 6    3ずつ増分
	// 	f.Println(rune_ja)  // 12354, 12356, 12358
	// }

	// i := 0
	// for {
	// 	f.Println(i)
	// 	i++
	// 	if i == 5 {
	// 		break  // 最内部(最も近い)forを抜ける
	// 	}
	// }

	// for i := 0; i < 10; i++ {
	// 	if (i % 2 == 1) {  // 1, 3, 5, 7, 9
	// 		if (i % 3 == 0) {  // 3, 9
	// 			// 奇数かつ３の倍数の場合はスキップ
	// 			continue
	// 		}
	// 	}
	// 	f.Println(i)  // 0〜9のうち、3と9以外を出力
	// }

	// f.Println("for文にネスト")

	// for num := 0; num < 3; num++ {  // 0,1,2で、３回まわる
	// 	for n := 0; n < 3; n++ {  // 0,1,2で、３回まわる
	// 		if n == 1 {
	// 			continue  // forがネストしている場合、最内部のforにおける現在のループをスキップし、次のループへと移る
	// 		}
	// 		f.Printf("n = %d\n", n)
	// 	}
	// 	f.Printf("num = %d\n", num)
	// }
	// n = 0
	// n = 2
	// num = 0
	// n = 0
	// n = 2
	// num = 1
	// n = 0
	// n = 2
	// num = 2

	// x := 30
	// if x == 1 {
	// 	f.Println("xは1です")
	// } else if x == 2 {
	// 	f.Println("xは2です")
	// } else if x == 3 {
	// 	f.Println("xは3です")
	// } else {
	// 	f.Println("xは1〜3以外です")
	// }

	// 条件式は必ず論理値(bool値)を返す必要がある
	// if (true) {
	// 	// 正常実行
	// 	f.Println("成立")
	// }

	// if (1) {  // non-bool 1 (type int) used as if condition
	// 	// コンパイルエラー
	// 	f.Println("成立")
	// }

	// 条件式の前に簡易文を挿入可能
	// 簡易文の評価(実行)がされてから、条件式の評価がされる
	// 下記の例では、変数x,yを定義してから、条件の評価がされる
	// 簡易文とは、式や代入分、暗黙な変数定義など複雑な構造を持たない単一の文のこと
	// if x, y := 1, 2; x < y {
	// 	f.Printf("x(%d) is less than y(%d)\n", x, y)
	// }

	// x, y := 3, 5
	// if n := x * y; n % 2 == 0 {
	// 	f.Println("偶数")
	// } else {
	// 	f.Println("奇数")
	// }
	// f.Println(n) // 変数nはifブロック内でのみ有効なので、未定義エラーが出る

	// 簡易文はエラー処理にも威力を発揮する
	// doSomething関数の２つ目の戻り値にエラーの有無を返すようにしておけば、エラー有の場合の処理を行うことができる
	// if _, err = doSomething(); err != nil {
	// 	f.Println("doSomething関数からエラーが返ってきました")
	// 	// エラー発生後の処理
	// }

	// for {
	// 	// 条件を指定しないと、無限ループになる
	// }

	// for i := 0; i < 10; i++ {
	// 	f.Println(i)
	// }

	// f.Println(foo.Hoge())

	// mathパッケージをパッケージ名なしでインポートしたことにより、パッケージ名の指定なしで参照可能になった
	// ただし、名前の重複には注意
	// f.Println(Pi)

	// foo.Init()

	// fmt.Println(foo.Max)
	// fmt.Println(foo.internal_const)  // エラー

	// fmt.Println(foo.FooFunc(5))
	// f.Println(foo.FooFunc(5))
	// fmt.Println(foo.internalFunc(5))  // エラー

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

	// fmt.Println(doSomething())  // 0

	// fmt.Println(returnMulti())  // 0 5

	// fmt.Println(ignoreArgs(100, 200))

	// fmt.Println(f(200, 300))

	// fmt.Printf("%T\n", func(x, y int) int { return x + y })  // func(int, int) int

	// f := func(x, y int) int { return x + y }

	// fmt.Println(f(200, 300))
}

// func doSomething() (a int) {
// 	return
// }

// 上記は下記の省略形

// func doSomething() int {
// 	var a int
// 	return a
// }

// func returnMulti() (x, y int) {
// 	y = 5
// 	return
// 	// 
// }

// 上記は下記の省略形

// func returnMulti() (int, int) {
// 	var x, y int
// 	y = 5
// 	return x, y
// }

// 戻り値を破棄することができるのと同様、引数を破棄することもできる
// func ignoreArgs(_, _ int) int {
// 	return 1
// }

// var f = func(x, y int) int { return x + y }
