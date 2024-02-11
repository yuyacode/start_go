package main

import (
	"fmt"
)

// panicとrecoverは、正確には構文ではない
// 定義済み関数として提供されている機能
// いわゆる「例外処理」に類するものとして考えて良い
// ただ、Goのランタイムを強制的に停止させる機能を持つため、決して多用すべきものではない

/*
------------------------
panic
------------------------
*/

// panicは定義済み関数として、次の形式で提供されている
// func panic(v interface{})

func panicMethod() {
	// 中断時までに登録されていたdeferは、全て実行される
	defer fmt.Println("on defer 1")
	defer fmt.Println("on defer 2")

	// panicを実行すると、即座にランタイムパニックが発生し、実行中の関数は中断する
	panic("runtime error!")  // ここでエラー終了
	
	// 出力は下記
	// panic: runtime error!    // どのような原因でランタイムが停止したのか、詳しい情報を付加するべき
	// goroutine 1 [running]:
	// main.t(...)
	//         /go/src/app/panic_recover.go:17
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


/*
------------------------
recover
------------------------
*/

// recoverとは、panicにより発生したランタイムパニックによるプログラム中断を回復するための機能
// recoverは、deferと組み合わせて使うのが原則
// panicは関数の実行を中断して、deferの評価に移行するので、実質的にdeferの中でしか動作しない

func recoverMethod() {
	defer func() {
		// 最後に実行されるdeferの中で、xがnilではない場合はパニックが実行されたと判断できるので、そこで色々やる
		if x := recover(); x != nil {
			// xは、panicに渡されたinterface{}
			fmt.Println("panic発生")
			fmt.Println(x)  // panic!!
		}
	}()
	panic("panic!!")

	// panicが発生した場合、実行されない
	// 上記のpanicをコメントアウトした場合は、当然走る
	fmt.Println("hello after panic")
}

func testRecoverWrapper() {
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
}

// panicとrecoverを組み合わせることによって、ある種の例外処理を実現可能
// しかし、よほどの場合を除いて使用する機会はないと考えていい
// そもそも関数がpanicを起こす可能性があることを前提にデザインされることは絶対に避けるべき
func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic発生 int=%v\n", v)
			case string:
				fmt.Printf("panic発生 string=%v\n", v)
			default:
				fmt.Println("panic発生 unknown")
			}
		}
	}()
	panic(src)
}