package main

import (
	"fmt"
	"regexp"
)

func regexpPkgTest1() {
	// regexp.MatchString関数は、実行されるたびに正規表現をコンパイルする
	// つまり、この関数が呼び出される度に、指定された正規表現パターンの解析と準備（コンパイル）が行われる
	// 正規表現のコンパイルは、パターンを解析し、それを実行するための内部表現に変換するプロセスのこと
	// このプロセスを通じて、正規表現エンジンはパターンマッチングを効率的に行えるようになる
	// しかし、MatchStringのような関数が毎回正規表現をコンパイルすることには欠点がある
	// 特に、同じ正規表現パターンを繰り返し使用する場合、このコンパイルプロセスが冗長になり、パフォーマンスに影響を与える可能性がある
	// 正規表現のコンパイルは、比較的コストの高い操作であるため、パフォーマンスが重要なアプリケーションではこの点を考慮する必要がある
	fmt.Println(regexp.MatchString("A", "ABC"))  // true <nil>    マッチしたかどうかbool、error
	fmt.Println(regexp.MatchString("A", "XYZ"))  // false <nil>
}

func regexpPkgTest2() {
	// regexp.Compileは、一度のコンパイルで済むため、効率的
	rePatternObj, err := regexp.Compile("A")
	if err != nil {
        fmt.Println("コンパイルエラー:", err)
        return
    }
	fmt.Printf("%T\n", rePatternObj)  // *regexp.Regexp
	fmt.Println(rePatternObj)  // A 
	fmt.Println(err)           // nil
	fmt.Println(rePatternObj.MatchString("ABC"))  // true
	fmt.Println(rePatternObj.MatchString("XYZ"))  // false
}