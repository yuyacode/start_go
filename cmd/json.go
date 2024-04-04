package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id int
	Name string
	Email string
	Created time.Time
}

func jsonPkgTest1() {
	u := new(User)
	u.Id = 1
	u.Name = "田中太郎"
	u.Email = "aaa@sample.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
	// {
	// 	"Id":1,
	// 	"Name":"田中太郎",
	// 	"Email":"aaa@sample.com",
	// 	"Created":"2024-04-04T02:11:08.602933089Z"
	// }
}

func jsonPkgTest2() {
	src := `
	{
		"Id":1,
		"Name":"田中太郎",
		"Email":"aaa@sample.com",
		"Created":"2024-04-04T02:11:08.602933089Z"
	}
	`

	u := new(User)
	err := json.Unmarshal([]byte(src), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u)
	// &{
	// 	Id:1 
	// 	Name:田中太郎 
	// 	Email:aaa@sample.com 
	// 	Created:2024-04-04 02:11:08.602933089 +0000 UTC
	// }
}