package main

import "fmt"

func main() {
	//右側の変数から型推論される
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

