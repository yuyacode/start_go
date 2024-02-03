// package foo; import ("fmt"); func main() {fmt.Println("Hello")};

// go fmt コマンドを使用すると、上記の１行を下記のように整形してくれる
// 整形したいファイルが存在するディレクトリに移動し、go fmtを打つだけ
// どのプロジェクトでもgo fmtコマンドに準拠したフォーマットにすることで、コーディングルールによる戦争が起きなくなる

package foo

import (
	"fmt"
)

func main() { fmt.Println("Hello") }
