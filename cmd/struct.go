package main

import (
	"fmt"
)

func struct_func() {

	// type [定義する型] [既存の型]
	// int型にMyIntというエイリアスを定義
	// typeは、既に定義されている型をもとに、新しい型を定義するための機能
	type MyInt int

	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1)  // 5
	fmt.Println(n2)  // 7
	// mainパッケージに属するMyInt型であることが示されている
	// Goでは、カスタム型（この場合はMyInt）を定義すると、その型は定義されたパッケージ名をプレフィックスとして持つ
	fmt.Printf("%T\n", n1)  // main.MyInt
	fmt.Printf("%T\n", n2)  // main.MyInt

	// 型のエイリアスを定義することの利点
	// map[string][2]float64のような複雑な型をAreaMapというエイリアスで定義することで、プログラムから複雑な型定義を取り除くことができる
	type (
		IntPair [2]int
		Strings []string
		AreaMap map[string][2]float64
		IntChannel chan []int
	)
	pair := IntPair{1, 2}
	strs := Strings{"AA", "BB", "CC"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntChannel)

}