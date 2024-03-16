package main

import (
	"fmt"
)

type point struct{X, Y int}

// レシーバーが値型
// レシーバーに値型が定義されたメソッドでは、呼び出し時にレシーバーそのもののコピーが発生する
// これにより、メソッドの呼び出し側とメソッド内部において、レシーバーの実体が異なる
func (p point) set(x, y int) {
	p.X = x
	p.Y = y
}

// レシーバーがポインタ型
func (p *point) setPointer(x, y int) {
	p.X = x
	p.Y = y
}

func method2Func() {
	p1 := point{}
	p1.set(1, 2)
	fmt.Println(p1.X)  // 0
	fmt.Println(p1.Y)  // 0

	p2 := &point{}
	p2.set(1, 2)
	fmt.Println(p2.X)  // 0
	fmt.Println(p2.Y)  // 0

	p3 := point{}
	p3.setPointer(1, 2)
	fmt.Println(p3.X)  // 1
	fmt.Println(p3.Y)  // 2

	p4 := &point{}
	p4.setPointer(1, 2)
	fmt.Println(p4.X)  // 1
	fmt.Println(p4.Y)  // 2

	// レシーバーが値型だとしてもポインタ型だとしても、呼び出し側は値型, ポインタ型の双方から呼び出すことができる
	// メソッドに対してレシーバーが値渡しされるか参照渡しされるかの違いは、レシーバー型が値型かポインタ型かによって決定する
}

// memo：構造体型のフィールドは全て非公開にし、全ての処理を公開されたメソッドを介して行うようにすることで、インターフェースが明快でメンテナンス性の高いパッケージを作り上げることができる