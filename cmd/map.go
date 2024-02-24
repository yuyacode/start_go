package main

import (
	"fmt"
)

func map_method() {
	var m map[int]string
	fmt.Println(m)  // map[]

	m1 := make(map[int]string)
	m1[1] = "US"
	m1[81] = "Japan"
	m1[86] = "China"
	fmt.Println(m1)  // map[1:US 81:Japan 86:China]
	
	m1[1] = "Korea"
	fmt.Println(m1)  // map[1:Korea 81:Japan 86:China]

	// 浮動小数点数の型なし定数をキーに使用する場合には、注意が必要
	// 下記３種類のキーは、float64型の精度に置き換えると全て1.0に自動で丸められてしまう
	m2 := make(map[float64]int)
	m2[1.000000000000000000000000001] = 1
	m2[1.000000000000000000000000002] = 2
	m2[1.000000000000000000000000003] = 3
	fmt.Println(m2)  // map[1:3]

	m3 := make(map[int]int)
	m3[1] = 1
	m3[2.0] = 2
	m3[3+0i] = 3
	fmt.Println(m3)  // map[1:1 2:2 3:3]
}
