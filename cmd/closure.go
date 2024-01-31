package main

import (
	"fmt"
)

func init() {
	f := later()

	fmt.Println(f("Golang"))   // ""
	fmt.Println(f("is"))       // "Golang"
	fmt.Println(f("awesome"))  // "is"
}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}