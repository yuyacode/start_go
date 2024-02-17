package main

import (
	"fmt"
)

func slice() {
	
	// int型のスライス
	// var s []int
	// var s []float64
	// var s []string
	// fmt.Println(s)  // []  上記全て、ゼロ値は[]らしい

	// 要素数と容量が10であるint型のスライスを生成
	// s1 := make([]int, 10)
	// fmt.Println(s1)  // [0 0 0 0 0 0 0 0 0 0]

	// s2 := make([]float64, 3)
	// fmt.Println(s2)  // [0 0 0]
	// s2[0] = 3.14
	// s2[1] = 6.28
	// fmt.Println(s2)  // [3.14, 6.28, 0]
	// fmt.Println(s2[4])  // ランタイムパニック

	// 要素数を調べる関数、len関数
	// s3 := make([]int, 8)
	// fmt.Println(len(s3))  // 8

	// len関数は、配列にも使用可能
	// s4 := [3]int {1, 2, 3}
	// fmt.Println(len(s4))  // 3

	// 容量を調べる関数、cap関数
	// s5 := make([]int, 5)
	// fmt.Println(len(s5))  // 5
	// fmt.Println(cap(s5))  // 5

	// s6 := make([]int, 5, 10)
	// fmt.Println(len(s6))  // 5
	// fmt.Println(cap(s6))  // 10

	// スライスはmake関数を使用せず、配列型のリテラルと同様の書き方で定義可能
	// この書き方の場合、容量は個別に指定できないため、次のコードの場合は要素数5, 容量5のスライスが出来上がる
	s7 := []int {1, 2, 3, 4, 5}
	fmt.Println(len(s7))  // 5
	fmt.Println(cap(s7))  // 5
}