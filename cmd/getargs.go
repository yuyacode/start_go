package main

import (
	"os"
	"fmt"
)

// $GOPATH(/go/src)に移動し、go build -o getargs app/getargs.goで、getargsという名前の実行ファイルをビルド
// ./getargs 123 456 789を実行
// 次の出力が得られる
// length = 4
// ./getargs
// 123
// 456
// 789

func main() {
	fmt.Printf("length = %d\n", len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}
}