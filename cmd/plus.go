package main

import (
	"fmt"
)

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello GO")
}

// 戻り値の型は１つに省略できないみたい。単にintと定義すると、単一の整数値のみを返す関数と認識されてしまうみたい
// (int, int)をint, intではコンパイルエラーになる。()は必須
func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}