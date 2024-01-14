package main

import (
	"fmt"
	"./animals"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}