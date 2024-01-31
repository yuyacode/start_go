package foo

import (
	// 
)

const (
	Max = 100
	internal_const = 1
)

func FooFunc(n int) int {
	return internalFunc(n)
}

func internalFunc(n int) int {
	return n + 1
}

// 日本語など大文字、小文字のない文字の場合、小文字と同様に他パッケージから参照できない