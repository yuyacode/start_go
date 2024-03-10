package main

import (
	"fmt"
)

// typeを使って定義された任意の型には、メソッドを定義することができる
// 元から存在する組み込み型（int, float64, stringなど）には、メソッドを定義することができない

type MyInt int

func methodFunc() {
	myInt1 := MyInt(3)
	isPositive := myInt1.IsPositive()
	fmt.Println(isPositive)  // true
}

func (myInt MyInt) IsPositive() bool {
	return myInt > 0
}