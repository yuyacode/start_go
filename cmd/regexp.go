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

func regexpPkgTest3() {
	// Compileを使用しても良いが、MustCompileを使用してコンパイルする方が望ましいかもしれない
	// MustCompileは、コンパイルエラーが発生した際に、エラーを返すのではなく、直接ランタイムパニックを発生させる
	// 正規表現のパターンは、固定であるケースがほとんどなので、テストさえしっかり行っていれば、コンパイルエラーが発生する可能性は少なく、つまりエラーを捕捉することによるメリットが少ない
	// エラーを返さないことで、エラーチェックのコードが省略され、プログラムの可読性が向上する
	// 正規表現のパターンを動的に組み立てる場合には、regexp.Compileを使用するメリットも考えられる
	// エラーハンドリングの挙動を除けば、操作や性能に違いはない
	regexp.MustCompile("ABC")
}