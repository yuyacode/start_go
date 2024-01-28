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

func main() {
	x, y := one()
	fmt.Printf("%d, %d\n", x, y)
	// 
}

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}