package main

import (
	"os"
	"fmt"
	"log"
	"io"
	"bufio"
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
	mainFileReadBytes, mainFileReadErr := mainFileHandler.Read(mainFileReadResult)  // Readは、*os.File型に紐づくメソッド
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

func osFunc3() {
	mainFileHandler, mainFileOpenErr := os.Open("app/app.go")
	if mainFileOpenErr != nil {
		log.Fatal(mainFileOpenErr)
	}
	defer mainFileHandler.Close()

	reader := bufio.NewReader(mainFileHandler)  // 変数readerに、バッファリングされたリーダー（Buffered Reader） = 入出力（I/O）操作を最適化するために使用される一時的なデータ保持領域（バッファ）を備えたリーダーを格納
	var output []byte
	buf := make([]byte, 128)

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		output = append(output, buf[:n]...)
		if err == io.EOF {
			break
		}
	}

	fmt.Println(string(output)) // osFunc2関数では、バイト数が足りていなかった関係で、ファイルの中身を全て読み取れていなかったが、ここでは全て読み取れた
}

func osFunc4() {
	mainFileHandler, mainFileOpenErr := os.Open("app/app.go")
	if mainFileOpenErr != nil {
		log.Fatal(mainFileOpenErr)
	}
	defer mainFileHandler.Close()

	mainFileReadResult := make([]byte, 128)
	mainFileReadBytes, mainFileReadErr := mainFileHandler.ReadAt(mainFileReadResult, 50)  // 50バイト目から
	if mainFileReadErr == io.EOF {  // ファイルの終わりに達した場合、エラーではなく正常な終了として扱う
		fmt.Println("Reached EOF - no error, but end of file")  // 出力される
		mainFileReadErr = nil  // EOFはエラーではないため、エラー変数をnilにリセット
	} else if mainFileReadErr != nil {
		log.Fatal(mainFileReadErr)
	}

	fmt.Println(mainFileReadResult)          // ファイルから読み込んだバイト列    [111 110 32 61 ...... 0 0 0 0 0]
	fmt.Println(string(mainFileReadResult))  // ファイルから読み込んだ文字列（バイト列を文字列へ変換した結果）    app.goの中身（途中から）
	fmt.Println(mainFileReadBytes)           // 読み込んだバイト数    108    
}

func osFunc5() {
	mainFileHandler, mainFileOpenErr := os.Open("app/app.go")
	if mainFileOpenErr != nil {
		log.Fatal(mainFileOpenErr)
	}
	defer mainFileHandler.Close()

	// ファイル内のシーク = ファイルの読み書き位置を移動する操作
	mainFileReadOffset1, _ := mainFileHandler.Seek(10, os.SEEK_SET)  // ファイルの先頭から10バイト目
	mainFileReadOffset2, _ := mainFileHandler.Seek(-2, os.SEEK_CUR)  // 現在のオフセットから-2バイト目
	mainFileReadOffset3, _ := mainFileHandler.Seek(0, os.SEEK_END)   // ファイルの末尾から0バイト目

	fmt.Println(mainFileReadOffset1)  // 10
	fmt.Println(mainFileReadOffset2)  // 8
	fmt.Println(mainFileReadOffset3)  // 158
}

func osFunc6() {
	mainFileHandler, mainFileOpenErr := os.Open("app/app.go")
	if mainFileOpenErr != nil {
		log.Fatal(mainFileOpenErr)
	}
	defer mainFileHandler.Close()

	mainFileInfo, mainFileStatErr := mainFileHandler.Stat()  // // os.FileInfo型のファイル情報を取得
	if mainFileStatErr != nil {
		log.Fatal(mainFileStatErr)
	}
	fmt.Println(mainFileInfo.Name())     //  app.go
	fmt.Println(mainFileInfo.Size())     //  158
	fmt.Println(mainFileInfo.Mode())     //  -rw-r--r--
	fmt.Println(mainFileInfo.ModTime())  //  2024-01-31 08:32:00.78771583 +0000 UTC 
	fmt.Println(mainFileInfo.IsDir())    //  false
}


func osFunc7() {
	createdFileHandler, _ := os.Create("app/osCreateTest.go")  // 新規ファイル作成  指定したファイルが既に存在する場合は、ファイルの内容が削除され、サイズ０のファイルに更新されることに注意
	createdFileInfo, _ := createdFileHandler.Stat()
	fmt.Println(createdFileInfo.Name())     //  osCreateTest.go
	fmt.Println(createdFileInfo.Size())     //  0
	fmt.Println(createdFileInfo.Mode())     //  -rw-r--r--
	fmt.Println(createdFileInfo.ModTime())  //  2024-03-24 05:08:46.591843193 +0000 UTC
	fmt.Println(createdFileInfo.IsDir())    //  false
}