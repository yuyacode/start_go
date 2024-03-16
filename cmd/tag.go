package main

import (
	"fmt"
	"reflect"
	"encoding/json"
)

// 構造体には、タグという、フィールドにメタ情報を付与する機能がある
// タグは文字列リテラル（"文字列"）か、RAW文字列リテラル（`RAW文字列`）を使用可能
// タグ内のダブルクオートが意味を持つ場面が多いため、RAW文字列リテラルを使用することが多い
type user struct {
	number int "番号"
	address string "住所"
	age uint "年齢"
}

// encoding/jsonパッケージは、構造体の公開フィールドにのみアクセスできる
// フィールドが非公開の場合、JSONの処理中に無視される
// JSONのマーシャリング/アンマーシャリングにおいて、型名の公開/非公開は関係ない
type student struct {
	Number   int    `json:"number"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
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

	// タグは、プログラム内で定義された構造体のフィールドに、文字列を使って柔軟性の高いメタ情報を追加する仕組み
	// タグの内容がどのように利用されるかは、ライブラリやプログラムの実装による

	// 構造体のタグが有効に活用される例として、jsonパッケージを使ったプログラムがある
	// jsonパッケージを利用すると、構造体型などのデータ構造をJSON形式のテキストに簡単に変換することができる
	// jsonパッケージは、与えられた構造体のフィールドのタグ内にjson:"［キー名］"という形式の文字列を見つけると、自動的にその情報を拾い出して、出力するJSONテキストのキー名として利用する
	// このように、タグはプログラムやライブラリに対してのメタ情報として機能する
	student1 := student{Number: 1, Nickname: "あああ", Gender: "男性"}
	bs, _ := json.Marshal(student1)
	fmt.Println(string(bs))  // {"number":1,"nickname":"あああ","gender":"男性"}
}