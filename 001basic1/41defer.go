package main

import "fmt"

func main() {
	//deferはreturnするまで評価を遅延させる
	defer fmt.Println("world")

	fmt.Println("hello")
}

