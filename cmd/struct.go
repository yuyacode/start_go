package main

import (
	"fmt"
)

// Callbackという明確な型名を定義することで、プログラムの可読性向上
// 複雑な型定義が煩わしい場合に、一考する価値がある
type Callback func(int) int

func struct_func() {

	// type [定義する型] [既存の型]
	// int型にMyIntというエイリアスを定義
	// typeは、既に定義されている型をもとに、新しい型を定義するための機能
	// 追記：typeによる型エイリアスの定義に関して、バージョンによっては関数内ではできないかもしれない。つまり、パッケージレベルでしか定義できないかも。
	//      正確な情報は現時点では分かっていないので、どちらの可能性も考慮しておく
	// type MyInt int

	// var n1 MyInt = 5
	// n2 := MyInt(7)
	// fmt.Println(n1)  // 5
	// fmt.Println(n2)  // 7
	// mainパッケージに属するMyInt型であることが示されている
	// Goでは、カスタム型（この場合はMyInt）を定義すると、その型は定義されたパッケージ名をプレフィックスとして持つ
	// fmt.Printf("%T\n", n1)  // main.MyInt
	// fmt.Printf("%T\n", n2)  // main.MyInt

	// 型のエイリアスを定義することの利点
	// map[string][2]float64のような複雑な型をAreaMapというエイリアスで定義することで、プログラムから複雑な型定義を取り除くことができる
	// type (
	// 	IntPair [2]int
	// 	Strings []string
	// 	AreaMap map[string][2]float64
	// 	IntChannel chan []int
	// )
	// pair := IntPair{1, 2}
	// strs := Strings{"AA", "BB", "CC"}
	// amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	// ich := make(IntChannel)

	// ここでtype定義すると、sum関数の箇所でundefined: Callbackが出る
	// ここで定義したtypeのスコープは、この関数内なので、別関数からは参照,アクセスできない
	// なので、ここではなくパッケージレベルで定義する必要がある
	// type Callback func(int) int
	// res := sum(
	// 	[]int {1, 2, 3, 4, 5},
	// 	func(i int) int {
	// 		return i * 2
	// 	},
	// )
	// fmt.Println(res)  // 30

	// T0型とint型には互換性がある
	// T1型とint型には互換性がある
	// しかし、同じintをベースにしたT0とT1の間には互換性はない
	// 同じ型から派生した場合であっても、エイリアス間では互換性がないことに注意
	// type T0 int
	// type T1 int

	// 「structで定義された構造体にtypeを使って新しい型名を与える」という順序
	// 型名が大文字で始まる場合、他パッケージからもアクセス可能な「公開」された型になる
	// 型名が小文字で始まる場合、定義されたパッケージ内部でのみアクセス可能な「非公開」な型となる
	// フィールドに関しても同様で、先頭が大文字は公開、先頭が小文字は非公開となる
	// 公開フィールドと非公開フィールドを１つの構造体内で混在させることは可能
	// 例え公開フィールドが定義されていたとしても、型自体が非公開な場合、公開フィールドに対して他パッケージからはアクセスできず、すなわち非公開として扱われる
	type Point struct {
		X int  // Xというint型のフィールド
		Y int  // Yというint型のフィールド
	}

	// 型が同じ場合は、下記のように一括定義も可能
	// type Point struct {
	// 	X, Y int
	// }

	// 構造体は値型の一種
	var pt Point
	fmt.Println(pt.X)  // 0  定義のみ行った場合、ゼロ値で初期化
	fmt.Println(pt.Y)  // 0  定義のみ行った場合、ゼロ値で初期化
	pt.X = 10
	pt.Y = 20
	fmt.Println(pt.X)  // 10
	fmt.Println(pt.Y)  // 20

	// リテラル（複合リテラル）でも定義可能
	// 下記の場合、順序が重要であり、構造体のフィールド定義順序に合わせる必要がある
	pt1 := Point{1, 2}
	fmt.Println(pt1.X)  // 1
	fmt.Println(pt1.Y)  // 2

	// 上記の場合、与える値の順序と構造体のフィールド定義順序を完全一致させる必要があり、
	// 新たな要件に対応するために既存の構造体フィールド定義に変更があった場合、対応関係がずれてしまうリスクがある
	// なので、対応関係がずれないように、構造体定義が変更となるたびにコードを修正する手間が発生する
	// なので、フィールドを明示的に指定して値を定義する複合リテラルが存在する
	// 特別な理由がない限りは、こちらの複合リテラルを使用するべき。プログラムの読解性が向上する
	pt2 := Point{Y: 2000, X: 1000}  // 順序を気にする必要がない
	fmt.Println(pt2.X)  // 1000
	fmt.Println(pt2.Y)  // 2000

	// フィールドを明示的に指定する複合リテラルを使用することで、一部のフィールドのみ値を指定可能
	pt3 := Point{Y: 888}
	fmt.Println(pt3.X)  // 0
	fmt.Println(pt3.Y)  // 888

	// フィールド名に日本語を用いることも一応できる
	// type Person struct {
	// 	ID int
	// 	部署 string
	// }
	// p := Person{ID: 17, 部署: "営業部"}

	// フィールド名は省略可能
	// 省略した場合、フィールド名=型名となる
	// intやstringなど基本的なものである場合は、使用するメリットはほとんどない
	type TT struct {
		int
		string
	}
	var tt TT
	fmt.Println(tt.int)     // 0
	fmt.Println(tt.string)  // ""

	// 無名フィールドを定義することもできる
	// フィールド名に_を与えると、そのフィールドは無名フィールドになる
	// 無名フィールドは、名前の通りフィールド名が存在しないので、参照や代入といった操作はできない
	// また同じ理由で、複合リテラルを使ったフィールド値の初期化もできない
	// しかし下記の出力結果から分かる通り、フィールドとしては間違いなく存在する
	// 無名フィールドは、構造体のメモリ上のアライメント調整のために用意されている
	// メモリのアライメント調整とは、データを配置するメモリを領域を細かく調整する仕組みのこと
	// 各データ（構造体の場合は各フィールド）を適したメモリ領域に配置することで、メモリアクセスを最適化し、パフォーマンスが向上する
	// 高度な仕組みであり、通常のGoプログラミングの範囲では理解しておく必要はない
	// このような仕組みがあることを認識しておく
	type T1 struct {
		N uint
		_ int16
		s []string
	}
	t1 := T1{
		N: 34,
		s: []string{"A", "B", "C"},
	}
	fmt.Println(t1)  // {34 0 [A B C]}

	// 構造体を含む構造体
	// 構造体に埋め込む構造体型に明示的にフィールド名を定義するパターン
	// type Animal struct {
	// 	Name string
	// 	Feed Feed  // ここでは「undefined: Feed」と出るので、定義順序は重要そう
	// }
	type Feed struct {
		Name string
		Amount uint
	}
	type Animal struct {
		Name string
		Feed Feed
	}
	monkey := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name: "Banana",
			Amount: 10,
		},
	}
	fmt.Println(monkey)              // {Monkey {Banana 10}}
	fmt.Println(monkey.Name)         // Monkey
	fmt.Println(monkey.Feed.Name)    // Banana
	fmt.Println(monkey.Feed.Amount)  // 10
	// fmt.Println(monkey.Amount)  // 10
	monkey.Feed.Amount = 20
	fmt.Println(monkey.Feed.Amount)  // 20

	// 構造体を含む構造体
	// 構造体に埋め込む構造体型にフィールド名を与えない（省略する）パターン
	type FeedFeed struct {
		Name string
		Amount uint
	}
	type AnimalAnimal struct {
		Name string
		FeedFeed  // フィールド名を省略  フィールド名は暗黙的にFeedFeedになる
	}
	elephant := AnimalAnimal{
		Name: "elephant",
		FeedFeed: FeedFeed{
			Name: "grass",
			Amount: 100,
		},
	}
	// 「フィールド名を省略して埋め込まれた構造体（FeedFeed）」のフィールド名（FeedFeed）が一意に定まる限り、中間のフィールド名は省略可能
	// 更に深くネストされた構造体の場合でも、同じことが通用する
	fmt.Println(elephant.Amount)  // 100
	elephant.Amount = 200
	fmt.Println(elephant.Amount)  // 200
	fmt.Println(elephant.FeedFeed.Amount)  // 200
	
}

func sum(ints []int, callbackFunc Callback) int {
	var sum int
	for _, v := range ints {
		sum += v
	}
	return callbackFunc(sum)
}