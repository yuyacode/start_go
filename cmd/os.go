package main

import (
	"os"
	"fmt"
	"log"
)

func osFunc() {
	defer func() {
		fmt.Println("defer")  // os.Exitによるプログラム停止を行うと、defer文は実行されない。破棄される
	}()
	fmt.Println("osパッケージ")
	os.Exit(1)  // exit status 1    終了ステータス１で終了
}

func osFunc1() {
	_, err := os.Open("notExistFile")
	if err != nil {
		log.Fatal(err)   // 2024/03/23 09:36:52 open notExistFile: no such file or directory
		                 // exit status 1
	}

	result, err := os.Open("foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)  // &{0x40000a6120}
}