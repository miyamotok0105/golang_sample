package main

import "fmt"

//struct (構造体)は、フィールド( field )の集まり
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

