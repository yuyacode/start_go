package main

import (
	"fmt"
)

func main() {
	// 無名関数を引数として受け取ることが可能
	process(func(s string) string {return "Hello" + " " + s + "!!!"}, "World")

	// 無名関数を戻り値として返すことが可能
	adder := createAdder(5)
	fmt.Println(adder(3))   // 8
	fmt.Println(adder(10))  // 15

	fmt.Printf("%#v\n", func(x, y int) int { return x + y })  // (func(int, int) int)(0x8c6c0)
	fmt.Printf("%#v\n", func(x, y int) int { return x + y } (200, 300))  // 500

	fmt.Println(plusAlias(100, 5))

	f := returnFunc()
	f()

	// 一度変数に入れずとも、直接実行することも可能
	returnFunc()()

	callFunction(func() { fmt.Println("関数を渡す") })
}

// 無名関数を引数として受け取ることが可能
func process(f func(string) string, value string) {
	fmt.Println(f(value))
}

// 無名関数を戻り値として返すことが可能
func createAdder(base int) func(int) int {
	return func(added int) int {
		return base + added
	}
}

func plus(x, y int) int {
	return x + y
}

var plusAlias = plus

func returnFunc() func() {
	return func() {
		fmt.Println("I am a function")
	}
}

func callFunction(f func()) {
	f()
}