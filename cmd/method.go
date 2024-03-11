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