package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//これは1個前のやつの通常の関数
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
