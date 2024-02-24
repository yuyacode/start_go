package main

import (
	"fmt"
)

func map_method() {
	// マップは要素の順序を保証しないので、同じマップでも出力するたびに要素の順序が変わる
	// ちなみに、配列やスライスは要素の順序を保証する
	// 配列やスライスは、インデックスで要素にアクセスするため
	// マップはキーで要素にアクセスするため、順序は保証されていない

	// var m map[int]string
	// fmt.Println(m)  // map[]

	// m1 := make(map[int]string)
	// m1[1] = "US"
	// m1[81] = "Japan"
	// m1[86] = "China"
	// fmt.Println(m1)  // map[1:US 81:Japan 86:China]
	
	// m1[1] = "Korea"
	// fmt.Println(m1)  // map[1:Korea 81:Japan 86:China]

	// 浮動小数点数の型なし定数をキーに使用する場合には、注意が必要
	// 下記３種類のキーは、float64型の精度に置き換えると全て1.0に自動で丸められてしまう
	// m2 := make(map[float64]int)
	// m2[1.000000000000000000000000001] = 1
	// m2[1.000000000000000000000000002] = 2
	// m2[1.000000000000000000000000003] = 3
	// fmt.Println(m2)  // map[1:3]

	// m3 := make(map[int]int)
	// m3[1] = 1
	// m3[2.0] = 2
	// m3[3+0i] = 3
	// fmt.Println(m3)  // map[1:1 2:2 3:3]

	// m4 := map[int]string {1: "Taro", 2: "Hanako", 3: "Jiro"}
	// fmt.Println(m4)  // map[1:Taro 2:Hanako 3:Jiro]

	// m5 := map[int]string {
	// 	1: "Taro",
	// 	2: "Hanako",
	// 	3: "Jiro",  // リテラルの継続を示すカンマを忘れずに
	// }
	// fmt.Println(m5)

	// マップの要素にスライスを指定することも可能
	// m6 := map[int][]int {
	// 	1: []int {1},
	// 	2: []int {1, 2},
	// 	3: []int {1, 2, 3},
	// }
	// fmt.Println(m6)  // map[1:[1] 2:[1 2] 3:[1 2 3]]

	// 下記のように、要素の[]int型のリテラルを描く場合の型は省略可能
	// map[int][]intの部分で、要素の型は[]intだと特定できるため
	// m7 := map[int][]int {
	// 	1: {1},
	// 	2: {1, 2},
	// 	3: {1, 2, 3},
	// }
	// fmt.Println(m7)

	// マップの要素に配列を指定可能だが、[...]intのように省略記法は使えない
	// そして、各要素の配列は、同サイズである必要がある
	// m8 := map[int][...]int {
	// 	1: [...]int {1, 2},
	// 	2: [...]int {1, 2, 3},
	// 	3: [...]int {1, 2, 3, 4},
	// }
	// fmt.Println(m8)

	// m9 := map[int][3]int {
	// 	1: [3]int {1, 2, 3},
	// 	2: [3]int {1, 2, 3},
	// 	3: [3]int {1, 2, 3},
	// }
	// fmt.Println(m9)

	// 要素の記述で用いるリテラルの型定義は省略可能
	// m10 := map[int][3]int {
	// 	1: {1, 2, 3},
	// 	2: {1, 2, 3},
	// 	3: {1, 2, 3},
	// }
	// fmt.Println(m10)

	// マップの要素の型がマップである場合
	// m11 := map[int]map[float64]string {
	// 	1: map[float64]string {3.14: "円周率"},
	// 	2: map[float64]string {3.14: "円周率"},
	// }
	// fmt.Println(m11)  // map[1:map[3.14:円周率] 2:map[3.14:円周率]]

	// こちらもリテラルの型表記は省略可能
	// m12 := map[int]map[float64]string {
	// 	1: {3.14: "円周率"},
	// 	2: {3.14: "円周率"},
	// }
	// fmt.Println(m12)  // map[1:map[3.14:円周率] 2:map[3.14:円周率]]

	// m13 := map[int]string {1: "A", 2: "B", 3: "C"}
	
	// s1 := m13[1]
	// fmt.Println(s1)  // "A"

	// キー9は存在しない、また要素はstring
	// これによって、stringのゼロ値である空文字が入る
	// 配列やスライスにおいては、存在しないキー（インデックス）を指定するとエラーになる
	// マップのみ特別で、初期値が入る
	// s2 := m13[9]
	// fmt.Println(s2)  // ""

	// マップ要素へのアクセスには、下記のような形式もある
	// １番目は値、２番目はキーに対応する要素が存在すればtrue、存在しなければfalse
	// ２番目の変数を使用することで、マップの要素への参照が成功したかどうか確認することができる
	// ２番目の変数名をokとするのは、Goにおける一種のイディオム
	// 特に事情がない限りは、okという変数名を設けるべき
	// m14 := map[int]string {1: "A", 2: "B", 3: "C"}
	// s, ok := m14[1]
	// fmt.Println(s, ok)  // A true
	// s1, ok1 := m14[9]
	// fmt.Println(s1, ok1)  // "" false
	// _, ok2 := m14[3]
	// fmt.Println(ok2)  // true

	// m15 := map[int]string {1: "A", 2: "B", 3: "C"}
	// if _, ok := m15[1]; ok {
	// 	fmt.Println("m15[1]は存在します")
	// } else {
	// 	fmt.Println("m15[1]は存在しません")
	// }

	// マップの要素型がスライスのような参照型である場合に、参照型変数の初期値がnilであることを利用して、
	// 下記のような「要素の存在有無チェック」が目的のコードは問題である
	// なぜなら、マップでは要素にnilを使用可能
	// つまり、意図してnilという要素を追加している可能性が否定できないため、nilだからといって、一概に要素が存在しないとは判断できない
	// 意図的にnilという値を要素に指定している可能性がある
	// m16 := map[int][]int {
	// 	1: nil,
	// 	2: {1, 2},
	// 	3: {1, 2, 3},
	// }
	// ok := m16[1]
	// if ok != nil {
	// 	fmt.Println("要素が存在")
	// } else {
	// 	fmt.Println("要素が存在しない")  // 意図的にnilという要素を追加しているものの、要素が存在しないケースとして扱われてしまう
	// }

	// なので、要素の存在チェックをしたい場合は、変数okを使用してチェックする方法が推奨である
	// if str, okok := m16[1]; okok {
	// 	fmt.Printf("要素は存在する。値は%s\n", str)  // 要素は存在する。値は[]  こっちが出力される
	// } else {
	// 	fmt.Printf("要素は存在しない。値は%s\n", str)
	// }

	// rangeによるfor
	// マップ内の要素の順序は保証されない点には注意
	m := map[int]string {
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}
	for k, v := range m {
		fmt.Printf("%d => %s\n", k, v)
	}



}
