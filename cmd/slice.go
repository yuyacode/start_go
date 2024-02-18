package main

import (
	"fmt"
)

func slice() {
	
	// int型のスライス
	// var s []int
	// var s []float64
	// var s []string
	// fmt.Println(s)  // []  上記全て、ゼロ値は[]らしい

	// 要素数と容量が10であるint型のスライスを生成
	// s1 := make([]int, 10)
	// fmt.Println(s1)  // [0 0 0 0 0 0 0 0 0 0]

	// s2 := make([]float64, 3)
	// fmt.Println(s2)  // [0 0 0]
	// s2[0] = 3.14
	// s2[1] = 6.28
	// fmt.Println(s2)  // [3.14, 6.28, 0]
	// fmt.Println(s2[4])  // ランタイムパニック

	// 要素数を調べる関数、len関数
	// s3 := make([]int, 8)
	// fmt.Println(len(s3))  // 8

	// len関数は、配列にも使用可能
	// s4 := [3]int {1, 2, 3}
	// fmt.Println(len(s4))  // 3

	// 容量を調べる関数、cap関数
	// s5 := make([]int, 5)
	// fmt.Println(len(s5))  // 5
	// fmt.Println(cap(s5))  // 5

	// s6 := make([]int, 5, 10)
	// fmt.Println(len(s6))  // 5
	// fmt.Println(cap(s6))  // 10

	// スライスはmake関数を使用せず、配列型のリテラルと同様の書き方で定義可能
	// この書き方の場合、容量は個別に指定できないため、次のコードの場合は要素数5, 容量5のスライスが出来上がる
	// s7 := []int {1, 2, 3, 4, 5}
	// fmt.Println(len(s7))  // 5
	// fmt.Println(cap(s7))  // 5

	// arr1 := [5]int {1, 2, 3, 4, 5}
	// arr1 := []int {1, 2, 3, 4, 5}  // スライスにも同様に使用可能
	// s8 := arr1[0:2]  // [1, 2]
	// s9 := arr1[2:]  // [3, 4, 5]
	// s10 := arr1[:4]  // [1, 2, 3, 4]
	// s11 := arr1[:]  // [1, 2, 3, 4, 5]
	// s98 := arr1[2:2]  // []
	// s99 := arr1[2:1]  // invalid slice index: 2 > 1

	// s12 := arr1[len(arr1)-2:]  // [4, 5]
	// fmt.Println(arr1[0:6])  // ランタイムパニック

	// s10[0] = 100
	// fmt.Println(s10)  // [100, 2, 3, 4]
	// fmt.Println(arr1)  // [100, 2, 3, 4, 5]  要素の変更は、元の配列に影響を及ぼす

	// s10 = append(s10, 500)
	// fmt.Println(cap(s10))  // 5
	// 容量以内での要素の追加であれば、それは元の配列にも影響を及ぼす
	// fmt.Println(s10)  // [100, 2, 3, 4, 500]
	// fmt.Println(arr1)  // [100, 2, 3, 4, 500]

	// 容量を超えた要素の追加の場合、元となる配列に要素を追加した新しい配列が作成され、その新しい配列を参照するスライスが新規作成される
	// これにより、新しいスライスは元の配列を参照しなくなるので、スライスへの要素の追加が、元の配列に影響を及ぼさなくなる
	// この挙動を「スライスの拡張」というらしい
	// s10 = append(s10, 600)
	// fmt.Println(s10)  // [100, 2, 3, 4, 500, 600]
	// fmt.Println(arr1)  // [100, 2, 3, 4, 500]

	// スライスに対して容量を超えた要素の追加が行われると、Goランタイムは元の配列よりも大きな新しい配列を割り当て、スライスに追加された要素を含むこの新しい配列を参照するようにスライスを更新します。
	// このプロセスを「スライスの拡張」と呼びます。

	// 文字列に対しても使用可能
	// str1 := "abcde"[1:3] // bc
	// fmt.Println(str1)

	// しかし、文字を単位にしている訳ではなく、バイト列が単位になっている
	// 文字列"あいうえお"は、UTF-8エンコードにおいて各文字に３バイトを要する
	// str2 := "あいうえお"[3:12] // いうえ
	// fmt.Println(str2)

	// appendを用いて、複数要素追加可能
	// s13 := []int {1, 2, 3}
	// s13 = append(s13, 4, 5, 6)
	// fmt.Println(s13)  // [1 2 3 4 5 6]
	
	// 既存のスライスs13を用いて、新しいスライスs14を作成可能
	// s14 := append(s13, 7, 8, 9)
	// fmt.Println(s14)  // [1 2 3 4 5 6 7 8 9]

	// 既存のスライスの組み合わせで新しいスライスを作成可能
	// ２つ目の引数の書き方が、arg...という特殊な形式になっている点には注意
	// arg...という特殊な書き方は、追加する要素を１つの集合としてまとめて第２引数に渡す場合に必要
	// 指定された集合内の各要素を個別の引数としてappend関数に渡す
	// これを「スライスの展開」という
	// s15 := append(s13, s14...)
	// fmt.Println(s15)  // [1 2 3 4 5 6 1 2 3 4 5 6 7 8 9]

	// []byte型のスライスと文字列型の組み合わせ時に、特殊な書き方が許容されている
	// var b []byte
	// b = append(b, "あいう"...)
	// b = append(b, "えお"...)
	// fmt.Println(b)  // [227 129 130 227 129 132 227 129 134 227 129 136 227 129 138]

	// 容量が不足すると、自動的に拡張するスライスの性質
	// 容量が不足したタイミングで、Goランタイムが容量を倍増させている
	// 容量の拡張が頻繁に発生する状況は、好ましい状況ではない
	// s16 := make([]int, 0, 0)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s16), cap(s16))  // 0 0
	// s16 = append(s16, 1)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s16), cap(s16))  // 1 1
	// s16 = append(s16, []int {2, 3, 4}...)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s16), cap(s16))  // 4 4
	// s16 = append(s16, 5)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s16), cap(s16))  // 5 8
	// s16 = append(s16, []int {6, 7, 8, 9}...)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s16), cap(s16))  // 9 16
	// s16 = append(s16, []int {10, 11, 12, 13}...)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s16), cap(s16))  // 13 16
	// s16 = append(s16, []int {14, 15, 16, 17}...)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s16), cap(s16))  // 17 32
	
	// 必ずしも倍増という訳ではなく、どの程度の容量が割り当てられるかは、Goランタイムに依存する
	// s17 := make([]int, 1024, 1024)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s17), cap(s17))  // 1024, 1024
	// s17 = append(s17, 0)
	// fmt.Printf("要素数：%d、容量：%d\n", len(s17), cap(s17))  // 1025, 1280

	// スライスにスライスの値を一括でコピーするための関数：copy関数
	// copy関数は、コピーした要素数を返す
	// srcをコピーして新しいメモリに割り当てるため、内容はコピーされるため同じであるものの、メモリが異なるためお互いを参照しない
	// すなわち完全に別物である
	// dst := make([]int, 4)
	// src := []int {1, 2, 3}
	// num := copy(dst, src)
	// fmt.Println(num, dst)  // 3 [1 2 3 0]

	// dstとscrのうち、要素数の小さい方に合わせる動きが行われる
	// dst1 := make([]int, 4)
	// src1 := []int {1, 2, 3, 4, 5, 6, 7}
	// num1 := copy(dst1, src1)
	// fmt.Println(num1, dst1)  // 4 [1, 2, 3, 4]

	// 文字列の場合、バイト単位でコピー
	// 下記の場合、"あいう"をコピー
	// dstBytes := make([]byte, 9)
	// srcString := "あいうえお"
	// copiedBytes := copy(dstBytes, srcString)
	// fmt.Println(copiedBytes, dstBytes)  // 9 [227 129 130 227 129 132 227 129 134]

	// 完全スライス式：３つのパラメータを取るスライス式
	// [low : high : max]
	// 0 <= low <= high <= max <= cap(元となる配列やスライス)
	// 簡易スライス式との違いは、maxの指定によってスライスの容量をコントロールできる
	// スターティングGo本の、完全スライス式の章に、容量の計算方法に関する分かりやすい図が載っている
	arrSample := [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrSample1 := arrSample[2:4]
	fmt.Printf("要素数：%d、容量：%d\n", len(arrSample1), cap(arrSample1))  // 2 8(元の配列の要素数10 - low2)
	arrSample2 := arrSample[2:4:4]
	fmt.Printf("要素数：%d、容量：%d\n", len(arrSample2), cap(arrSample2))  // 2 2(max4 - low2)
	// arrSample3 := arrSample[2:4:3]
	// fmt.Printf("要素数：%d、容量：%d\n", len(arrSample3), cap(arrSample3))  // invalid slice index: 4 > 3
	arrSample4 := arrSample[2:4:6]
	fmt.Printf("要素数：%d、容量：%d\n", len(arrSample4), cap(arrSample4))  // 2 4(max6 - low2)

}