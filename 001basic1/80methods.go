package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
//メソッドは、レシーバ引数を伴う関数
//Goには、クラス( class )のしくみはありませんが、型にメソッド( method )を定義できます。
//v Vertexってのがレシーバ引数に入ってる

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	//vにAbs関数が入ってる
	fmt.Println(v.Abs())
}

