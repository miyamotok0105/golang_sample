package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

//rangeはスライスを反復する
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

