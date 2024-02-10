package main

import (
	"fmt"
)

func init() {
	ints := integers()

	fmt.Println(ints())  // 1
	fmt.Println(ints())  // 2
	fmt.Println(ints())  // 3

	// クロージャを別に新しく生成
	// クロージャ間で状態は共有されない
	otherInts := integers()
	fmt.Println(otherInts())  // 1
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}