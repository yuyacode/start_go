package main

import (
	"fmt"
)

// 予約後deferを使って、関数の終了時に実行する関数を登録することができる
// deferで登録可能なのは関数の呼び出しだけであり、変数定義や計算式を登録することはできない
// 変数定義などをしたい場合は、それを関数にラップし、その関数を呼び出す代替策が考えられる
func runDefer() {
	defer fmt.Println("defer")
	fmt.Println("done")
}

func runDeferMulti() {
	// deferの登録が複数ある場合、後で登録されたdeferから順に実行される
	// 下記の場合、3, 2, 1の順で出力される
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("done")
}

// deferを使用することで、リソースの解放やファイルのクローズ,ロック解除など、後片付け処理を実行する際に役立つ
func fileOperation() {
	file, err := os.Open("path/to/file")
	if err != nil {
		return
	}
	defer file.Close()
}

// defer文に複数の処理を登録したい場合は、無名関数も利用可能
// しかし、deferに登録できるのは関数呼び出しだけなので、無名関数の後に()を記載し、実行まで保証することが必要
func deferClosure() {
	// 下記の場合は、A,B,Cの順で出力される
	// deferの登録自体は１つしかない
	defer func() {
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	}()
	fmt.Println("deferClosure関数")
}