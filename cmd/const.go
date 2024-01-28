package main

import (
	"fmt"
)

const ONE = 1

// 変数同様、まとめて定義可能
// const (
// 	X = 1
// 	Y = 2
// 	Z = 3
// )

// 定数の値を省略することができる
const (
	X = 1  // X = 1
	Y      // Y = 1
	Z      // Z = 1
	S1 = "あ"  // S1 = "あ"
	S2         // S2 = "あ"
)

func main() {
	x, y := one()
	fmt.Printf("%d, %d\n", x, y)
	// 
}

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}