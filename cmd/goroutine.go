package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go文は、並行処理を司る特別な機能
// go文は、defer文と同様に、関数呼び出し形式の式を受け取る

// 下記２関数を実行すると、それぞれの無限ループが並行して実行され、各ループで出力する文字列が不規則に表示される
func parent() {
	go child()
	for {
		fmt.Println("parent loop")
	}
}
func child() {
	for {
		fmt.Println("child loop")
	}
}

// 無名関数を使用して、上記を下記のように書くことも可能
func parentAlias() {
	go func() {
		for {
			fmt.Println("child loop")
		}
	}()
	for {
		fmt.Println("parent loop")
	}
}

// Goは、スレッドよりも小さい処理単位である「ゴルーチン(goroutine)」が並行して動作するように設計されている
// go文は、goroutineを新規作成して、そのgoroutine(新たに並行処理するもの)をスケジュールし、ランタイムに追加するための機能


// runtimeパッケージには、Goのランタイム自身の情報を参照したり、動作をコントロールするための機能が含まれる
func goroutineRuntimeTest() {
	go fmt.Println("syncパッケージは未使用")                      // これ(goroutine)を追加することで、NumGoroutineは2になる
	fmt.Printf("NumCPU : %d\n", runtime.NumCPU())              // 8  // Goランタイムが動作する環境のCPUの数。使用できるCPUコア数
	fmt.Printf("NumGoroutine : %d\n", runtime.NumGoroutine())  // 1  // Goランタイム上で動作しているゴルーチンの数
	fmt.Printf("Version : %s\n", runtime.Version())
}

// 上記goroutineRuntimeTest関数では、メインゴルーチン(goroutineRuntimeTest関数内の処理)と新たに生成されたゴルーチン(go文)の2つのゴルーチンが並行処理されるが、
// メインゴルーチンが早期に終了する関係で、go文によるゴルーチンは実行されず終了する(go文によるゴルーチンの実行が間に合わない)
// それを防ぐため、メインゴルーチンが新しいゴルーチンの完了を待つようにする必要がある。下記に再現
func goroutineRuntimeTestUsedSync() {

	// sync.WaitGroupは、複数のゴルーチンが完了するのを待つために使用される同期化プリミティブ(原型)の一つ
	// これを使用することで、一つのゴルーチンが他複数のゴルーチンの実行全てが終わるまで待機することができる
	var wg sync.WaitGroup  // syncパッケージのWaitGroup型の変数を定義。WaitGroupのインスタンスと考えてもいいかも
	wg.Add(1)  // WaitGroupのカウンタを指定された値だけ増やす。新しいゴルーチンを起動する前に、そのゴルーチンの数だけカウンタを増やす

	go func() {
		defer wg.Done()  // ゴルーチンがその処理を終了したことをWaitGroupに通知する。内部的にはWaitGroupのカウンタを1減らす。これは通常、ゴルーチンの最後にdefer文と共に使用される
		fmt.Println("goroutineの生成と追加")
	}()

	wg.Wait()  // WaitGroupのカウンタが0になるまで、以降の処理の実行をブロックする。つまり、全てのゴルーチンがDone()を呼び出して終了することを待つ

	fmt.Println("syncパッケージを使用")
	fmt.Printf("NumCPU : %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine : %d\n", runtime.NumGoroutine())
	fmt.Printf("Version : %s\n", runtime.Version())
}