package main

import (
	"os"
	"fmt"
)

func osFunc() {
	defer func() {
		fmt.Println("defer")  // os.Exitによるプログラム停止を行うと、defer文は実行されない。破棄される
	}()
	fmt.Println("osパッケージ")
	os.Exit(1)  // exit status 1    終了ステータス１で終了
}