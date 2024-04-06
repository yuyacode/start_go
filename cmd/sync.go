package main

import (
	"fmt"
	"time"
	"sync"
)

var st struct{A, B, C int}

func UpdateAndPrint(n int) {
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st.A, st.B, st.C)
}

func syncTest1() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	wg.Wait()

	// 出力
	// ３つのゴルーチンが１つの共有変数stに対して操作を行っているため、A,B,C全てのフィールドが同じ値で出力されないケースが存在する
	// フィールドCの値を書き換えている間に、他ゴルーチンがフィールドAやBを書き換えており、結果、出力時にはフィールドAやBの値が自身が代入した値ではない値で出力されている
	// 0 0 0
	// 0 0 0
	// 1 0 0
	// 1 1 1
	// 2 1 1
	// 2 2 1
	// 2 2 2
	// 2 2 2
	// 2 2 2
}