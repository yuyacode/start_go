package main

import (
	"fmt"
)

// typeを使って定義された任意の型には、メソッドを定義することができる
// 元から存在する組み込み型（int, float64, stringなど）には、メソッドを定義することができない

type MyInt int

func methodFunc() {
	myInt1 := MyInt(3)
	isPositive := myInt1.IsPositive()
	fmt.Println(isPositive)  // true

	// 匿名構造体（typeキーワードを使用せずに直接定義された構造体）に対して、直接メソッドを定義することはできない
	// メソッドを定義するためには、そのメソッドが属する型が名前を持っている必要がある
	// 匿名構造体に対してメソッドを定義しようとすると、コンパイルエラーになる
	person := struct{Name string}{Name: "太郎"}
	person.Greet()  // person.Greet undefined (type struct { Name string } has no field or method Greet)
}

func (myInt MyInt) IsPositive() bool {
	return myInt > 0
}

func (person struct{Name string}) Greet() string {  // invalid receiver type struct { Name string } (struct { Name string } is not a defined type)
	return "Hello, my name is " + person.Name
}