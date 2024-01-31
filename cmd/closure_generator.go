package main

import (
	"fmt"
)

func init() {
	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	// クロージャを別に新しく生成
	// クロージャ間で状態は共有されない
	otherInts := integers()
	fmt.Println(otherInts())
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}