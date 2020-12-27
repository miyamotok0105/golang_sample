package main

import "fmt"

//これらのスライス式は等価です:
//var a [10]int
// a[0:10]
// a[:10]
// a[0:]
// a[:]

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

