package main

import "fmt"

//フィボナッチ関数の実装をしてみる

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

