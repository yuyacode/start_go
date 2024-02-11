package main

import (
	"fmt"
)

// panicは、正確には構文ではない
// 定義済み関数として提供されている機能
// いわゆる「例外処理」に類するものとして考えて良い
// ただ、Goのランタイムを強制的に停止させる機能を持つため、決して多用すべきものではない

// panicは定義済み関数として、次の形式で提供されている
// func panic(v interface{})

func panicMethod() {
	// 中断時までに登録されていたdeferは、全て実行される
	defer fmt.Println("on defer 1")
	defer fmt.Println("on defer 2")

	// panicを実行すると、即座にランタイムパニックが発生し、実行中の関数は中断する
	panic("runtime error!")  // ここでエラー終了
	
	// 出力は下記
	// どのような原因でランタイムが停止したのか、詳しい情報を付加するべき
	// panic: runtime error!
	// goroutine 1 [running]:
	// main.t(...)
	//         /go/src/app/panic.go:17
	// main.main()
	//         /go/src/app/main.go:13 +0x38
	// exit status 2
	
	fmt.Println("Hello")  // 実行されない

	// 中断時以降に登録されているdeferは、全て実行されない
	defer fmt.Println("on defer 3")
	defer fmt.Println("on defer 4")
}

// panicは、プログラムにおいて「これ以上処理を継続しようがない状態」「回復不能な状態」を表すために使用される
// 決して、一般的なエラー処理で使用するものではない