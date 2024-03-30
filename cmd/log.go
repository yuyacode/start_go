package main

import (
	"fmt"
	"log"
	"os"
)

func logPkgTest1() {
	log.Print("ログ1行目\n")
	log.Println("ログ2行目")
	log.Printf("ログ%d行目\n", 3)

	fmt.Println(os.Stdout)  // &{0x40000b2060}

	// ログを任意ファイルに出力する
	f, err := os.Create("log/test1.log")
	if err != nil {
		return
	}
	fmt.Printf("%T\n", f)  // *os.File
	log.SetOutput(f)
	log.Println("log message")
}