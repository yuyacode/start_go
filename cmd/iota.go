package main

import (
	// 
)

// iotaは0から始まり、式として使用されるたびに1ずつ増分していく性質を持つ
// iotaを参照して得られる値は、constが構成するブロックごとに0にリセットされる
// 定数を定義する場合は、1つのconstブロックに無理に詰め込むことはせず、関連性を持つ定数をグルーピングした方が良い

const (
	a = iota  // 0
	b = iota  // 1
	c = iota  // 2
)

const (
	// 下記でもOK
	d = iota  // 0
	e         // 1
	f         // 2
)

const (
	// iotaが生成する値を0からではなく1から始めたい場合は、下記のように書く
	g = 1 + iota  // 1
	h             // 2
	i             // 3
)

// iotaの使用を途中で遮断した場合
// iotaは参照の有無によらず、constブロックの中で定数が定義されるたびに1ずつ増加する
// 定数群に付与されたインデックスとしてイメージする方が適切かも
const (
	// 
	j = iota  // 0
	k         // 1
	l         // 2
	m = 17    // 17
	n = iota  // 4
	o         // 5
)

// iotaを途中から定義した場合
// こちらも同様、インデックスのようなイメージで通用する
// 前提として、iotaの使用,未使用が混在する定数定義は、好ましい書き方ではない
const (
	p = 999   // 999
	q = iota  // 1
	r         // 2
)

func main() {
	// 
}
