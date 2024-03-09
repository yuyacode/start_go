package main

import (
	"fmt"
)

// Callbackという明確な型名を定義することで、プログラムの可読性向上
// 複雑な型定義が煩わしい場合に、一考する価値がある
type Callback func(int) int

func struct_func() {

	// type [定義する型] [既存の型]
	// int型にMyIntというエイリアスを定義
	// typeは、既に定義されている型をもとに、新しい型を定義するための機能
	// 追記：typeによる型エイリアスの定義に関して、バージョンによっては関数内ではできないかもしれない。つまり、パッケージレベルでしか定義できないかも。
	//      正確な情報は現時点では分かっていないので、どちらの可能性も考慮しておく
	// type MyInt int

	// var n1 MyInt = 5
	// n2 := MyInt(7)
	// fmt.Println(n1)  // 5
	// fmt.Println(n2)  // 7
	// mainパッケージに属するMyInt型であることが示されている
	// Goでは、カスタム型（この場合はMyInt）を定義すると、その型は定義されたパッケージ名をプレフィックスとして持つ
	// fmt.Printf("%T\n", n1)  // main.MyInt
	// fmt.Printf("%T\n", n2)  // main.MyInt

	// 型のエイリアスを定義することの利点
	// map[string][2]float64のような複雑な型をAreaMapというエイリアスで定義することで、プログラムから複雑な型定義を取り除くことができる
	// type (
	// 	IntPair [2]int
	// 	Strings []string
	// 	AreaMap map[string][2]float64
	// 	IntChannel chan []int
	// )
	// pair := IntPair{1, 2}
	// strs := Strings{"AA", "BB", "CC"}
	// amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	// ich := make(IntChannel)

	// ここでtype定義すると、sum関数の箇所でundefined: Callbackが出る
	// ここで定義したtypeのスコープは、この関数内なので、別関数からは参照,アクセスできない
	// なので、ここではなくパッケージレベルで定義する必要がある
	// type Callback func(int) int
	// res := sum(
	// 	[]int {1, 2, 3, 4, 5},
	// 	func(i int) int {
	// 		return i * 2
	// 	},
	// )
	// fmt.Println(res)  // 30

	// T0型とint型には互換性がある
	// T1型とint型には互換性がある
	// しかし、同じintをベースにしたT0とT1の間には互換性はない
	// 同じ型から派生した場合であっても、エイリアス間では互換性がないことに注意
	// type T0 int
	// type T1 int

	// 「structで定義された構造体に、typeを使って新しい型名を与える」という順序
	type Point struct {
		X int  // Xというint型のフィールド
		Y int  // Yというint型のフィールド
	}

	// 型が同じ場合は、下記のように一括定義も可能
	type Point struct {
		X, Y int
	}

}

func sum(ints []int, callbackFunc Callback) int {
	var sum int
	for _, v := range ints {
		sum += v
	}
	return callbackFunc(sum)
}