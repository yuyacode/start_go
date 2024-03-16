package main

import (
	"fmt"
)

// typeを使って定義された任意の型には、メソッドを定義することができる
// 元から存在する組み込み型（int, float64, stringなど）には、メソッドを定義することができない

type MyInt int
type student struct {
	number int
	name string
	selfPr string
}
type Strings []string
type User struct {
	Id int
	Name string
}
type Point struct {
	X, Y int
}

func methodFunc() {
	myInt1 := MyInt(3)
	isPositive := myInt1.IsPositive()
	fmt.Println(isPositive)  // true

	// 匿名構造体（typeキーワードを使用せずに直接定義された構造体）に対して、直接メソッドを定義することはできない
	// メソッドを定義するには、そのメソッドが属する型が名前を持っている必要がある
	// 匿名構造体に対してメソッドを定義しようとすると、コンパイルエラーになる
	// person := struct{Name string}{Name: "太郎"}
	// person.Greet()  // person.Greet undefined (type struct { Name string } has no field or method Greet)

	student1 := student{number: 1, name: "藍沢", selfPr: "誰よりも強い男です。負けません"}
	student1.showOverview()
	student2 := student{number: 2, name: "井上", selfPr: "誰よりも頭がいい男です。負けません"}
	student2.showOverview()

	fmt.Println(Strings{"A", "B", "C"}.join(","))  // A,B,C

	fmt.Println(NewUser(1, "太郎"))  // &{1 太郎}

	// メソッドを関数型として参照するときには、「レシーバーの型.メソッド」のように書くことができる
	f := (*Point).ToString  // メソッドを関数型として参照    イメージは、〇〇型に属する〇〇メソッドを変数に代入
	f(&Point{X: 7, Y: 11})  // [7, 11]  // レシーバーを経由せず、通常の関数として呼び出すことが可能    変数に代入された関数を呼び出す、このとき第１引数にはレシーバーとして扱われる情報を渡してあげる

	// 複雑なコードだが、変数を使わずに「メソッドを関数型として参照」→「関数の呼び出し」までをまとめて書くと次のようになる
	((*Point).ToString)(&Point{X: 11, Y: 13})  // [11, 13]

	// メソッドを関数型として参照する場合、レシーバー.メソッド名という書き方も存在する
	pHoge := &Point{X: 22, Y: 8}  // イメージ：レシーバーを変数に格納
	fHoge := pHoge.ToString  // イメージ：変数pHogeの型に紐づくメソッドToStringを参照し、fHogeに格納
	fmt.Println(fHoge())  // [22, 8]    // イメージ：関数fHogeを実行。このときレシーバーとして扱われる情報は既に定まっているので、引数で渡す必要なし
}

// メソッドはレシーバーを第１引数として受け取る単なる関数
func (p *Point) ToString() string {
	return fmt.Sprintf("[%d, %d]", p.X, p.Y)
	// fmt.Println(fmt.Sprintf("[%d, %d]", p.X, p.Y))  // Sprintfは整形までが担当なので、出力したい場合はこのようにする
}

// Goには、オブジェクト指向言語に見られるコンストラクタ機能はないが、慣例的に「型のコンストラクタ」というパターンを利用する
// 型のコンストラクタを表す関数は、「New型名」のように命名するのが一般的
// 型のコンストラクタは、対象の型のポインタ型を返すように定義するのが望ましい。ただ別に値型を返しても全然良い
// 型のコンストラクタをパッケージ内部でのみ利用するのであれば、newUserのように先頭を小文字にして非公開にするのが良い
// NewXXXという一種のイディオムは、Goの標準パッケージでも頻繁に利用されている
func NewUser(id int, name string) *User {
	u := new(User)  // User型のポインタ型を生成
	u.Id = id
	u.Name = name
	return u
}

func (s Strings) join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}

func (myInt MyInt) IsPositive() bool {
	return myInt > 0
}

// func (person struct{Name string}) Greet() string {  // invalid receiver type struct { Name string } (struct { Name string } is not a defined type)
// 	return "Hello, my name is " + person.Name
// }

// 同じ型のインスタンスに対して行う操作をここでまとめることができる
func (student student) showOverview() {
	fmt.Println("概要")
	fmt.Printf("出席番号：%d\n", student.number)
	fmt.Printf("名前：%s\n", student.name)
	student.detail()  // レシーバーを通して、そのインスタンスの他メソッド（同じ型に紐づく他メソッド）にアクセスすることができる
}

func (student student) detail() {
	fmt.Println("詳細")
	fmt.Println(student.selfPr)
}