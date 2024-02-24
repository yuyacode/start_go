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

	m4 := map[int]string {1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m4)  // map[1:Taro 2:Hanako 3:Jiro]

	m5 := map[int]string {
		1: "Taro",
		2: "Hanako",
		3: "Jiro",  // リテラルの継続を示すカンマを忘れずに
	}
	fmt.Println(m5)

	// マップの要素にスライスを指定することも可能
	m6 := map[int][]int {
		1: []int {1},
		2: []int {1, 2},
		3: []int {1, 2, 3},
	}
	fmt.Println(m6)  // map[1:[1] 2:[1 2] 3:[1 2 3]]

	// 下記のように、要素の[]int型のリテラルを描く場合の型は省略可能
	// map[int][]intの部分で、要素の型は[]intだと特定できるため
	m7 := map[int][]int {
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	fmt.Println(m7)

	// マップの要素に配列を指定可能だが、[...]intのように省略記法は使えない
	// そして、各要素の配列は、同サイズである必要がある
	// m8 := map[int][...]int {
	// 	1: [...]int {1, 2},
	// 	2: [...]int {1, 2, 3},
	// 	3: [...]int {1, 2, 3, 4},
	// }
	// fmt.Println(m8)

	m9 := map[int][3]int {
		1: [3]int {1, 2, 3},
		2: [3]int {1, 2, 3},
		3: [3]int {1, 2, 3},
	}
	fmt.Println(m9)

	// 要素の記述で用いるリテラルの型定義は省略可能
	m10 := map[int][3]int {
		1: {1, 2, 3},
		2: {1, 2, 3},
		3: {1, 2, 3},
	}
	fmt.Println(m10)

	// マップの要素の型がマップである場合
	m11 := map[int]map[float64]string {
		1: map[float64]string {3.14: "円周率"},
		2: map[float64]string {3.14: "円周率"},
	}
	fmt.Println(m11)  // map[1:map[3.14:円周率] 2:map[3.14:円周率]]

	// こちらもリテラルの型表記は省略可能
	m12 := map[int]map[float64]string {
		1: {3.14: "円周率"},
		2: {3.14: "円周率"},
	}
	fmt.Println(m12)  // map[1:map[3.14:円周率] 2:map[3.14:円周率]]


}
