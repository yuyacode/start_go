package main

import (
	"bufio"
	"fmt"
	"strings"
)

func bufioPkgTest1() {
	s := `
	XXX
	YYY
	ZZZ
	`

	// strings.NewReader関数は、与えられた文字列を元にしたio.Readerインターフェースと互換性があるstrings.Reader型を生成
	// 「io.Readerインターフェースと互換性がある」とは、つまりそのインターフェースを実装しているということ
	// 従って、strings.NewReaderによって生成されるstrings.Readerオブジェクトは、io.Readerインターフェースを満たしているため、io.Readerを要求するあらゆる関数やメソッドに渡すことができる。これにより、文字列から読み取りを行う際に非常に便利なツールとなっている
	r := strings.NewReader(s)

	// スキャナの生成
	scanner := bufio.NewScanner(r)
	fmt.Printf("%T\n", scanner)  // *bufio.Scanner

	scanner.Scan()
	fmt.Println(scanner.Text())  // ""
	scanner.Scan()
	fmt.Println(scanner.Text())  // XXX
	scanner.Scan()
	fmt.Println(scanner.Text())  // YYY
	scanner.Scan()
	fmt.Println(scanner.Text())  // ZZZ
}