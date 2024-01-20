package main

import (
	// "fmt"
)

var n = 1

func increment() int {
	n = n + 1
	return n
}

func decrement() int {
	n = n - 1
	return n
}

func local_increment() int {
	var i int
	i = 100
	i = i + 100
	return i
}

func local_decrement() int {
	var i int
	i = 100
	i = i - 100
	return i
}