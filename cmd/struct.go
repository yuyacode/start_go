package main

import (
	"fmt"
)

// Callbackという明確な型名を定義することで、プログラムの可読性向上
// 複雑な型定義が煩わしい場合に、一考する価値がある
type Callback func(int) int

type multiInt struct {
	X, Y int
}

type user struct {
	id int
	name string
}

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

	// 「名前つきフィールドを使った組み込み」と「無名フィールドを用いた埋め込み」では、アクセス方法以外に「型の振る舞い」や「メソッドの継承」、「インターフェースの実装」などで違いがあるみたい
	// 正直、現時点ではアクセス方法以外に違いが分かっていないので、後々勉強が必要。下記はイメージを持つためのメモ
	// 名前つきフィールドを使った組み込みでは、「has-a」の関係を示唆する（AはBaseを持っている）
	// 無名フィールドを用いた埋め込みは、しばしば「is-a」の関係を示唆する（AはBaseである）

	// フィールド名を省略した埋め込み構造体では、「異なる構造体型に共通の性質を持たせる」局面で有効に利用することができる
	// 異なる構造体間で共有し得るフィールド群を、Base型のように別の構造体型として切り出して共通化することができる
	type Base struct {
		Id int
		Owner string
	}
	type A struct {
		Base
		Name string
		Area string
	}
	type B struct {
		Base
		Title string
		Bodies []string
	}
	aaa := A{
		Base: Base{Id: 17, Owner: "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}
	bbb := B{
		Base: Base{Id: 81, Owner: "Sho"},
		Title: "no title",
		Bodies: []string{"A", "B"},
	}
	fmt.Println(aaa.Id)     // 17
	fmt.Println(aaa.Owner)  // Taro
	fmt.Println(bbb.Id)     // 81
	fmt.Println(bbb.Owner)  // sho

	// ポインタ型の修飾子やパッケージのプレフィックス部分は無視され、純粋な型名部分が暗黙的なフィールド名として利用される
	// struct {
	// 	T1     // フィールド名はT1
	// 	*T2    // フィールド名はT2
	// 	P.T3   // フィールド名はT3
	// 	*P.T4  // フィールド名はT4
	// }

	// 構造体のフィールドに自身の型を含むパターンは、再帰的な定義であり、コンパイルエラー
	// type T struct {
	// 	T
	// }

	// フィールドに含まれる構造体型（T1）が、自身の型（T0）を含むようなパターンは、再帰的な定義であり、コンパイルエラー
	// type T0 struct {
	// 	T1
	// }
	// type T1 struct {
	// 	T0
	// }

	// 構造体型は主にtypeと組み合わせて利用するが、struct{}という構造体型そのものを型として利用することも可能
	// ただ、あえてこのような煩雑な定義を選択する必要はない
	s := struct{X, Y int}{X: 1, Y: 2}
	showStruct(s)  // {1, 2}

	// typeによって定義した構造体型のエイリアスは、元の構造体と互換性がある
	// なので、下記のようにIntPair型の変数valをshowStruct関数に渡しても、正常に動作する
	type IntPair struct {
		X, Y int
	}
	val := IntPair{X:3, Y:8}
	showStruct(val)  // {3 8}

	// 構造体は値型
	// 関数の引数で渡す際には、コピーが発生する
	aaaVal := AAA{X: "XXX", Y: "YYY"}
	swap(aaaVal)
	fmt.Println(aaaVal)  // {XXX YYY}  入れ替わっていない  2
	
	// ポインタを渡すように変更
	// 実はポインタ型を使用することが一般的で、値型として処理する局面は限定されている
	swap_P(&aaaVal)
	fmt.Println(aaaVal)  // {YYY XXX}  入れ替わっている  4

	// 指定した型のポインタ型を生成するための組み込み関数new
	type Person struct {
		Id int
		Name string
		Area string
	}
	person := new(Person)
	fmt.Println(person)  // &{0 "" ""}

	// 組み込み関数newは、構造体型限定ではなく、基本型や参照型にも使用可能
	// しかし、そのような使い方にはあまりメリットがなく、やはり主な用途は構造体型のポインタ生成のために利用される
	iii := new(int)  // この時点でiii自体とiiiが指すメモリアドレスの２メモリ領域が確保される
	fmt.Println(iii)  // 0x40000ae0c0
	fmt.Println(*iii)  // 0
	strSlice := new([]string)
	fmt.Println(strSlice)  // &[]
	fmt.Println(*strSlice  == nil)  // true
	// strSliceは何も参照していないというよりかは、メモリアドレスを保持している以上、特定メモリ領域を参照しているが、
	// 参照先にはデータが存在しない,空であり、実際のデータを参照している訳ではない（つまり、どこも参照していない）と認識されて、nilと判断されているみたいなイメージ

	// 組み込み関数newを使った構造体型のポインタ生成と、アドレス演算子&を伴った複合リテラルによる構造体型のポインタ生成の間に、動作上ほとんど違いはない
	personPointer1 := new(Person)
	personPointer1.Name = "太郎"
	personPointer1.Area = "渋谷"
	fmt.Println(personPointer1)  // &{0, "太郎", "渋谷"}
	fmt.Println(*personPointer1)  // {0, "太郎", "渋谷"}

	personPointer2 := &Person{Name: "三郎", Area: "シリコンバレー"}
	fmt.Println(personPointer2)  // &{0, "三郎", "シリコンバレー"}
	fmt.Println(*personPointer2)  // {0, "三郎", "シリコンバレー"}

	mi := &multiInt{X: 5, Y: 12}

	// 型に定義されたメソッドは、レシーバー.メソッドという形式で呼び出す
	mi.Render()

	// 構造体型をマップのキーや値に指定する場合、リテラル内で構造体型の型名を省略可能。通常はuser{id: ~~~, name: ~~~}と書かなければいけないが
	// キーが構造体のマップ
	map1 := map[user]string{
		{id: 1, name: "Taro"}: "Tokyo",
		{id: 2, name: "Jiro"}: "Fukuoka",
	}
	// 値が構造体のマップ
	map2 := map[int]user{
		1: {id: 1, name: "Taro"},
		2: {id: 2, name: "Jiro"},
	}
	fmt.Println(map1)  // map[{1 Taro}:Tokyo {2 Jiro}:Fukuoka]
	fmt.Println(map2)  // map[1:{1 Taro} 2:{2 Jiro}]

	// キーや値にスライスやマップを指定する場合も同様に型の省略が可能
	ms := map[int][]string {
		1: {"A", "B", "C"}
	}
	mm := map[int]map[int]string {
		1: {1: "one", 2: "two"}
	}

}

// メソッドとは、任意,特定の型に特化した,紐づいた関数のこと
// 型が違えばデータ構造が違う、データ構造が違えば行う（行うことができる）処理も違う、だからこそメソッドとは特定の型に特化した関数
// 各型にはそれぞれ独自のデータ構造があり、その構造に応じて適切な処理を行う必要がある。メソッドを使用することで、型ごとに特有の操作を定義し、それらの操作を簡単に再利用し、組み合わせることが可能になる

// 構造体はフィールドの集合体であり、メソッドはその構造体のインスタンスに対する操作を定義する。メソッドとは構造体というある種のデータに対する操作をまとめたもの
// メソッドに処理をまとめる利点
// ・再利用性：メソッドを通じて、特定の操作を繰り返し利用することができる。これにより、コードの重複を減らし、保守性が高まる
// ・分離：データと操作を分離することで、コードの簡潔性や可読性が向上する
// ・カプセル化：データ自体を直接変更するのではなく、メソッドを経由し、メソッドからのみデータ変更を行うようにする。これにより、データが一定のルールや条件に従って変更されることを保証する

// メソッド定義では、関数とは異なり、funcとメソッド名の間に「レシーバー（Receiver）」の変数名と型が必要
// 下記の場合、*multiInt型の変数miがレシーバーを指す
// メソッドにおけるレシーバーは、メソッドが操作する対象のインスタンスを指す。また、メソッドがそのレシーバーの型のインスタンスから呼び出されることを意味する
// レシーバーを通して、メソッドはそのインスタンスのフィールドや他メソッドにアクセス可能
// レシーバーは、メソッドがどの型に属しているか定義する。これにより、同じ名前のメソッドが異なる型で異なる振る舞いをすることが可能

// メソッドは関数同様、任意の引数と戻り値を定義することができる
func (mi *multiInt) Render() {
	fmt.Printf("<%d, %d>\n", mi.X, mi.Y)  // <5, 12>
}

type AAA struct {
	X, Y string
}

func swap(aaaVal AAA) {
	aaaVal.X, aaaVal.Y = aaaVal.Y, aaaVal.X  // 入れ替える
	fmt.Println(aaaVal)  // {YYY XXX}  ここでは入れ替わっている  1
}

func swap_P(aaaVal *AAA) {
	aaaVal.X, aaaVal.Y = aaaVal.Y, aaaVal.X  // 入れ替える
	fmt.Println(aaaVal)  // &{YYY XXX}  入れ替わっている  3
}

func showStruct(s struct{X, Y int}) {
	fmt.Println(s)
}

func sum(ints []int, callbackFunc Callback) int {
	var sum int
	for _, v := range ints {
		sum += v
	}
	return callbackFunc(sum)
}