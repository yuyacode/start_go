package main

import (
	"fmt"
	"log"
	"net/url"
)

func netUrlPkgTest1() {
	// URLを表す文字列をパースして、url.URL型の構造体を生成する
	u, err := url.Parse("https://sample.com/search?a=1&b=2#top")
	if err != nil {
		log.Fatal(err)
	}

	// url.URL型は、URLを構成する各要素をフィールドに持つ
	fmt.Println(u.Scheme)   // https
	fmt.Println(u.Host)     // sample.com
	fmt.Println(u.Path)     // /search
	fmt.Println(u.RawQuery) // a=1&b=2
	fmt.Println(u.Fragment) // top
	fmt.Println(u.IsAbs())  // true              // urlが絶対か相対か識別するメソッド
	fmt.Println(u.Query())  // map[a:[1] b:[2]]  // urlのクエリーの内容をマップとして参照するメソッド
}