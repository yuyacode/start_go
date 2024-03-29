package main

import (
	"fmt"
	"sync"
	"time"
)

func chan_func() {

	// チャネルは、ゴルーチンとゴルーチンの間でデータの受け渡しを司るためにデザインされた、Go特有のデータ構造
	// チャネルとは、複数のゴルーチン間で安全にデータを共有するための仕組み
	// つまり、ゴルーチンによる非同期処理を必要としないプログラムでは、原則使用する必要はない

	// int型のチャネルを表す型
	// var ch chan int

	// チャネルには、Goの他データ型にはない特殊なサブタイプを指定可能
	// <-chanを使用すると、そのチャネルは受信専用チャネルとなる
	// var ch1 <-chan int
	// chan<-を使用すると、送信専用チャネルとなる
	// var ch2 chan<- int
	// このようなオプションを指定しないchanは、送受信可能な双方向チャネルとして機能する

	// Goのデータ型は厳密であり、異なる型の変数同士の代入は原則コンパイルエラーになる
	// しかし、チャネルのサブタイプは少し事情が異なる
	// var (
	// 	ch3 chan int
	// 	ch4 <-chan int
	// 	ch5 chan<- int
	// )
	// 下記は、チャネルとそのサブタイプ間における代入可否
	// 送受信可能な双方向チャネルのみ、受信専用,送信専用に代入可能。それ以外はNG
	// ch4 = ch3  // OK
	// ch5 = ch3  // OK
	// ch3 = ch4  // NG
	// ch5 = ch4  // NG
	// ch3 = ch5  // NG
	// ch4 = ch5  // NG

	// チャネル自身は、本質的には受信も送信も可能なキューのようなデータ構造
	// 実際は多くの場面で、受信専用として処理されるか、送信専用として処理されるか、明確に分かれる
	// 基本となるのは、あくまでchanではあるものの、「局面に応じて<-chanやchan<-に切り替えることを意図した仕組み」であることを理解する

	// チャネルに関しても、組み込み関数makeを使用して生成可能
	// makeへ２番目の引数を渡すことで、チャネルのバッファサイズを指定可能
	// 明示的に指定されていない場合、バッファサイズは0
	// ch6 := make(chan int)     // バッファサイズ：0
	// ch7 := make(chan int, 8)  // バッファサイズ：8

	// バッファなしチャネルを使用する目的の一つは、ゴルーチン間での同期を実現すること
	// この同期により、送信と受信が一対一で対応し、送信されたデータが受信されるまでの間、送信側は次の操作に進めないため、同期処理が可能になる

	// チャネルとは、キュー（待ち行列）の性質を備えるデータ構造
	// チャネルのバッファとは、このキューを格納する領域のこと
	// すなわち、キューのサイズと見なすことができる

	// キューにはFIFOの性質があり、先にキューにenqueueされたデータを、先にdequeueできるという特性がある
	// すなわち、データを取り出す順序が保証されるという特性があり、Goのチャネルも同様の性質を備えている

	// チャネルが保持するデータに対する操作は「送信」か「受信」の２パターンのみ
	// 送受信ともに演算子<-を使用
	// ch8 := make(chan int, 10)

	// チャネルに整数５を送信
	// ch8 <- 5
	// チャネルから整数値を受信
	// i := <-ch8

	// チャネルはキューとしての性質を持つデータ構造ではあるが、Goプログラム中で単なるキューとして利用するようにはデザインされていない
	// チャネルはあくまでも複数のゴルーチン間で安全にデータを共有するための仕組みである

	// チャネルからデータを受信する処理は、裏返せば、他ゴルーチンがチャネルにデータを送信するのを待つということ
	// 唯一存在するメインゴルーチンは、チャネルからのデータ受信のために眠ってしまったものの、そのチャネルにデータを送信してくれるゴルーチンは存在しない
	// このような状態をGoランタイムはデッドロックであると検知し、ランタイムパニックを検知させたというのが事の真相
	// ch9 := make(chan int)
	// fmt.Println(<-ch9)  // fatal error: all goroutines are asleep - deadlock!

	// バッファなしチャネルを使用しているため、送信と受信が１対１の関係性になり、すなわち送信側と受信側で同期処理が行われる
	// これにより、「送信処理だけが早期に終了し、これによりメインゴルーチンが終了し、それによりサブゴルーチンが強制終了することで、受信側が全ての受信を終える前に処理が終了してしまう」事態を防ぐことができている
	// バッファありチャネルの場合は、これが起こるケースがあるので、sync.WaitGroupを使用してサブゴルーチンの終了を待ったり、チャネルを閉じることが必要（チャネルのクローズは後述）
	// ch10 := make(chan int)
	// go receiver(ch10)
	// i := 0
	// for i <= 100 {
	// 	ch10 <- i
	// 	i++
	// }

	// len関数：チャネルのバッファ内に溜められているデータ数を取得する
	// ch11 := make(chan string, 3)
	// ch11 <- "Apple"
	// fmt.Println(len(ch11))  // 1
	// ch11 <- "Banana"
	// ch11 <- "Cherry"
	// fmt.Println(len(ch11))  // 3

	// 次のようなコードは書くべきではない
	// 仮に、len(ch11) > 0が成立したとしても、次の瞬間にはチャネルの状態が変化している（バッファ内が空になっている）可能性が十分にあるため
	// if len(ch11) > 0 {
	// 	// バッファ内のデータが存在する場合
	// 	i := <-ch11
	// }

	// チャネルにおけるデッドロックとは、プログラム内のすべてのゴルーチンが、チャネル操作（送信または受信）でブロックされ、それ以上の進行が不可能になる状態を指す
	// この状態では、どのゴルーチンも待っている操作（送信または受信）を完了することができず、プログラムは処理を進めることができない

	// cap関数：チャネルのバッファサイズを取得
	// チャネルのバッファサイズは変更不可能であり、チャネル定義時のみ指定可能
	// なので、有効に使用できる局面は正直少ない
	// ch12 := make(chan string)
	// fmt.Println(cap(ch12))  // 0

	// ch13 := make(chan string, 3)
	// fmt.Println(cap(ch13))  // 3

	// クローズしたチャネルに送信することはできない
	// ch14 := make(chan int, 1)
	// close(ch14)
	// ch14 <- 1  // panic: send on closed channel

	// チャネルがクローズされても、チャネルのバッファ内に溜められたデータについては問題なく受信できる
	// バッファ内が空になった場合は、チャネルが内包する型の初期値を受信し続ける。ランタイムパニック等は発生しない
	// チャネルのクローズは、送信側の役割であり、受信側がチャネルをクローズすることはないと認識して良い
	// ch15 := make(chan int, 3)
	// ch15 <- 1
	// ch15 <- 2
	// ch15 <- 3
	// close(ch15)
	// fmt.Println(<-ch15)  // 1
	// fmt.Println(<-ch15)  // 2
	// fmt.Println(<-ch15)  // 3
	// fmt.Println(<-ch15)  // 0
	// fmt.Println(<-ch15)  // 0

	// 下記コードの場合、送信側がサブゴルーチンとしての立ち位置を取るため、sync.WaitGroupを使用した制御が不要になる
	// ch16 := make(chan int)
	// go produce(ch16)
	// consume(ch16)

	// sync.WaitGroup使って複数ゴルーチンを制御した書き方
	// var wg sync.WaitGroup
	// ch17 := make(chan int)
	// wg.Add(1)
	// go produceWg(ch17, &wg)  // $wgで、変数wgのメモリアドレスを渡す  呼び出し元と先で同じwgを共有しないと、Add,Done,Waitの状態を共有できない
	// wg.Add(1)
	// go consumeWg(ch17, &wg)
	// wg.Wait()

	// ch18 := make(chan int, 3)
	// ch18 <- 1
	// ch18 <- 2
	// ch18 <- 3
	// close(ch18)

	// チャネルがクローズされているかどうかは、少し不正確
	// 厳密には「チャネルのバッファ内が空であり、かつクローズされている状態」の場合に、変数okはfalseになる
	// var(
	// 	i int
	// 	ok bool
	// )
	// i, ok = <-ch18
	// fmt.Println(i, "", ok)  // 1 true
	// i, ok = <-ch18
	// fmt.Println(i, "", ok)  // 2 true
	// i, ok = <-ch18
	// fmt.Println(i, "", ok)  // 3 true
	// i, ok = <-ch18
	// fmt.Println(i, "", ok)  // 0 false

	// ch19 := make(chan int, 20)
	// go receive("1st goroutine", ch19)
	// go receive("2nd goroutine", ch19)
	// go receive("3rd goroutine", ch19)
	// i := 0
	// for i < 30 {
	// 	ch19 <- i
	// 	i++
	// }
	// close(ch19)
	// time.Sleep(3 * time.Second)  // だいぶ雑な処理だが、３つのサブゴルーチンの完了のために３秒待つ

	// 下記コードの場合、ch21のチャネルからデータを受信できないとゴルーチンは停止する
	// つまり、ch22のチャネルからの受信にいつまで経っても辿りつかない
	// Go言語において、チャネルからのデータ受信操作はブロッキング操作
	// e1 := <-ch21
	// e2 := <-ch22

	// selectステートメントを使用することで、複数チャネルに対する受信,送信処理に対して、ゴルーチンを停止させることなくコントロール可能
	// select {
	// case e1 := <-ch21 :
	// 	// ch21からの受信が成功した場合の処理
	// case e2 := <-ch22 :
	// 	// ch22からの受信が成功した場合の処理
	// default :
	// 	// case節が成立しなかった場合の処理
	// }

	// 当たり前だが、case節の式はチャネル操作を扱っている必要がある
	// 具体的は、受信処理、送信処理、２つのチャネル間の受信と送信を直接繋ぐ処理のいずれかである必要がある
	// case節の式の種類は、下記４種類
	// select {
	// case e1 := <-ch21 :
	// case e2, ok := <-ch2 :
	// case ch3 <- e3 :
	// case ch4 <- (<-ch5) :  // ch5から受信したデータをch4に送信
	// }

	// ch23 := make(chan int, 1)
	// ch24 := make(chan int, 1)
	// ch25 := make(chan int, 1)
	// ch23 <- 1
	// ch24 <- 2
	// 複数のcase節が成立する場合は、実行するcaseがランダムで選択される
	// 当然、１つのcaseしか成立しない場合は、そのcaseが実行される
	// select {
	// case <-ch23 :
	// 	fmt.Println("ch23から受信")
	// case <-ch24 :
	// 	fmt.Println("ch24から受信")
	// case ch25 <- 3 :
	// 	fmt.Println("ch25へ送信")
	// default :
	// 	fmt.Println("ここへは到達しない")
	// }

	channel1 := make(chan int)
	channel2 := make(chan int)

	go sendNumbers(channel1, 100 * time.Millisecond, 5)  // time.Millisecond  timeパッケージのMillisecondという定数で、１ミリ秒を返す
	go sendNumbers(channel2, 150 * time.Millisecond, 5)

	// チャネルを参照している変数にnilを代入することで、その変数はチャネルへの参照を失うことになる
	// つまり、チャネルへの操作を一時的に無効化する手段として使用される
	// selectステートメントにおいて、caseで扱っているチャネル変数がnilを参照していることを検知すると、そのcaseの評価すら実施されない
	// これにより、他のcaseの評価へと素早く移動でき、処理の効率が高まる。特定のゴルーチンからのデータ受信が完了した後に、残りのゴルーチンからのデータ受信に集中したい場合などに有用
	for {
		select {
		case num, ok := <-channel1 :
			if !ok {
				channel1 = nil
				fmt.Println("Channel1 is closed")
			} else {
				fmt.Println("Received from channel1：", num)
			}
		case num, ok := <-channel2 :
			if !ok {
				channel2 = nil
				fmt.Println("Channel2 is closed")
			} else {
				fmt.Println("Received from channel2：", num)
			}
		}
		if channel1 == nil && channel2 == nil {
			break
		}
	}
	fmt.Println("Finished receiving!")
	// Received from channel2： 0
	// Received from channel1： 0
	// Received from channel1： 1
	// Received from channel2： 1
	// Received from channel1： 2
	// Received from channel1： 3
	// Received from channel2： 2
	// Received from channel1： 4
	// Received from channel2： 3
	// Channel1 is closed
	// Received from channel2： 4
	// Channel2 is closed
	// Finished receiving!

}

func sendNumbers(channel chan int, delay time.Duration, max int) {
	for i := 0; i < max; i++ {
		channel <- i
		time.Sleep(delay)
	}
	close(channel)
}

func receive(name string, ch19 <-chan int) {
	for {
		i, ok := <-ch19
		if ok == false {
			break
		}
		fmt.Println(name, "：", i)
	}
	fmt.Println(name, " 完了")
}

func produceWg(ch17 chan<- int, wg *sync.WaitGroup) {  // wg *sync.WaitGroup  sync.WaitGroup型のポインタをwgとして受け取る
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch17 <- i
		time.Sleep(time.Second)
	}
	close(ch17)
}

func consumeWg(ch17 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// for i := range chという書き方は、チャネルからの受信を継続的に続け、受信したデータを処理しつつ、チャネルのクローズを検知すると自動でループを終了する特別な構文
	// チャネルからのデータの受信、処理、そしてチャネルのクローズ検知を簡潔に記述することができる
	// 明示的なクローズ検知のコードを書く必要がなくなる
	// 追記：もしかしたらクローズ検知はしてくれないかも。本にはそう書いてあった。どっちが正解か分からないので、一旦放置
	for i := range ch17 {
		fmt.Println("受信データ：", i)
	}
}

// func produce(ch16 chan<- int) {
// 	for i := 0; i <= 10; i++ {
// 		ch16 <- i
// 		time.Sleep(time.Second)  // ここで1秒待機させてみる
// 	}
// 	close(ch16)
// }

// func consume(ch16 <-chan int) {
// 	for {
// 		i, ok := <-ch16
// 		if !ok {
// 			fmt.Println("チャネルがクローズされました")
// 			break
// 		}
// 		fmt.Println("受信データ：", i)
// 	}
// 	// 
// }

// func receiver(ch10 <-chan int) {
// 	for {
// 		i := <-ch10
// 		fmt.Println(i)
// 	}
// }