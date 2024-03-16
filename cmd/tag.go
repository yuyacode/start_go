package main

import (
	"fmt"
	"reflect"
)

// 構造体には、タグという、フィールドにメタ情報を付与する機能がある
// タグは文字列リテラル（"文字列"）か、RAW文字列リテラル（`RAW文字列`）を使用可能
// タグ内のダブルクオートが意味を持つ場面が多いため、RAW文字列リテラルを使用することが多い
type user struct {
	number int "番号"
	address string "住所"
	age uint "年齢"
}

func tagFunc() {
	// プログラム内から構造体に付与されたタグを参照するプログラム
	// ここでは、プログラム実行時にタグの内容を動的に参照できることを理解する
	// 補足、プログラム実行時の型情報を参照するためにreflectパッケージを利用している
	user1 := user{number: 1, address: "神奈川県", age: 23}
	t := reflect.TypeOf(user1)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
		// number 番号
		// address 住所
		// age 年齢
	}
}