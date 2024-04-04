package main

import (
	"fmt"
	"io/ioutil"  // 1.16以降では非推奨になったパッケージ。ioutilパッケージの多くの関数がosやioパッケージに移動した
	"log"
	"net/http"
	"net/url"
)

func netHttpPkgTest1() {
	// これだけでGETメソッドを実行できる
	// エラーが発生しなければ、有効な*http.Response型によるレスポンスを得られる
	res, err1 := http.Get("https://google.com/")
	if err1 != nil {
		log.Fatal(err1)
	}

	// http.Response型のフィールド
	fmt.Println(res.StatusCode)  // 200
	fmt.Println(res.Proto)  // HTTP/2.0

	// レスポンスヘッダの参照
	fmt.Println(res.Header)  // map型のでかいデータが出力される
	fmt.Println(res.Header["Date"])  // [Thu, 04 Apr 2024 04:12:44 GMT]
	fmt.Println(res.Header["Content-Type"])  // [text/html; charset=ISO-8859-1]
	
	// リクエスト情報
	fmt.Printf("%+v\n", res.Request)
	// &{Method:GET URL:https://www.google.com/ Proto: ProtoMajor:0 ProtoMinor:0 Header:map[Referer:[https://google.com/]] Body:<nil> GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host: Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr: RequestURI: TLS:<nil> Cancel:<nil> Response:0x4000014120 ctx:0x40000b0010}
	fmt.Println(res.Request.Method)  // GET
	fmt.Println(res.Request.URL)  // https://www.google.com/

	// レスポンスボディの読み込み
	defer res.Body.Close()
	body, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(string(body))  // リクエスト先URL（https://google.com/）のviewのソースコードが出力される
}

func netHttpPkgTest2() {
	vs := url.Values{}  // url.Values{}は、URLのクエリパラメータを保持するためのマップを新しく作成。url.Valuesはmap[string][]string型のエイリアス
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	vs.Encode()  // url.Values型に紐づくメソッドEncodeにより、クエリ文字列にエンコードする

	res, err := http.PostForm("https://exsample.com/comments/post", vs)  // http.PostForm関数は、指定されたURLに対して、application/x-www-form-urlencoded形式のフォームデータ（ここではvsで指定したデータ）を含むHTTP POSTリクエストを送信
	// resは、GETメソッド同様、http.Response型

	if err != nil {
		log.Fatal(err)
	}
}

func netHttpPkgTest3() {
	f, err := os.Open("foo.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	res, err := http.Post("https://example.com/upload", "image/jpg", f)  // ファイルアップロード
	if err != nil {
		log.Fatal(err)
	}
	// Goの公式ドキュメントによれば、クライアントがレスポンスボディを読み終えた後、または必要ない場合でも、このボディ（ストリーム）をクローズする必要がある。
	// これは、ネットワークリソースを適切に解放し、接続を再利用するため。接続を再利用できるようにすることは、特に多くのリクエストを行うアプリケーションにおいて、パフォーマンスの向上につながる
	// レスポンスボディをクローズしないと、その接続がハングアップ（フリーズ）してしまい、リソースリークを引き起こす可能性がある。長期間実行されるアプリケーションでは、これが原因でファイルディスクリプタの枯渇やメモリリークなどの問題が発生する可能性がある
	// つまり、受け取ったレスポンスに対して何か操作を加えていなくても、クローズは必要
	defer res.Body.Close()

	// レスポンスの受け取りを破棄する場合でも、http.PostなどのHTTPリクエストから返されるhttp.ResponseオブジェクトのBodyをクローズすることは必要
	// たとえレスポンスボディを使わない、あるいはレスポンス自体を無視する（_を使用して受け取りを破棄する）場合でも、クローズ操作は行うべき
	// レスポンスボディをクローズしない場合、そのHTTP接続は再利用されずに放置されることになる。これはリソースの無駄遣いであり、プログラムが長時間実行される場合は、リソースリークにつながる可能性がある
	_, err := http.Post("https://example.com/upload", "image/jpg", f)
}