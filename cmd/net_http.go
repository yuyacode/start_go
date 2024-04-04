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