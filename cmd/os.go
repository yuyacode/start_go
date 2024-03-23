package main

import (
	"os"
	"fmt"
	"log"
	"io"
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

func osFunc2() {
	mainFileHandler, mainFileOpenErr := os.Open("app/app.go")
	if mainFileOpenErr != nil {
		log.Fatal(mainFileOpenErr)
	}
	defer mainFileHandler.Close()

	mainFileReadResult := make([]byte, 128)
	mainFileReadBytes, mainFileReadErr := mainFileHandler.Read(mainFileReadResult)
	if mainFileReadErr == io.EOF {  // ファイルの終わりに達した場合、エラーではなく正常な終了として扱う
		fmt.Println("Reached EOF - no error, but end of file")
		mainFileReadErr = nil  // EOFはエラーではないため、エラー変数をnilにリセット
	} else if mainFileReadErr != nil {
		log.Fatal(mainFileReadErr)
	}

	fmt.Println(mainFileReadResult)          // ファイルから読み込んだバイト列    [112 97 99 107 97 ...]
	fmt.Println(string(mainFileReadResult))  // ファイルから読み込んだ文字列（バイト列を文字列へ変換した結果）    app.goの中身
	fmt.Println(mainFileReadBytes)           // 読み込んだバイト数    128
}