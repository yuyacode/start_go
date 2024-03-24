package main

import (
	"fmt"
	"time"
)

func timePkgTest1() {
	now := time.Now()  // 現在時刻を取得
	fmt.Println(now)  // 2024-03-24 07:47:32.980856387 +0000 UTC m=+0.000018168

	createdTime := time.Date(2024, 3, 24, 16, 49, 40, 0, time.Local)  // 指定した日時を表すtime.Time型を自由に生成可能
	fmt.Println(createdTime)  // 2024-03-24 16:49:40 +0000 UTC

	// ある種のタイムアウト処理を実現
	// ５秒以内にチャネルchからの受信が行われず、１つ目のcaseが成立しなかった場合、５秒後にtime.After関数により現在時刻が送られてくるため、それを受信することで２つ目のcaseが成立し、selectステートメントを抜けるという、ある種のタイムアウト処理を実現
	select {
	case m := <-ch :
		fmt.Println(m)
	case <-time.After(5 * time.Second) :
		fmt.Println("time out")
	}

}