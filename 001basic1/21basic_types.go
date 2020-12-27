package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)

	num1   int     = 1
	num2   int8     = 2

	num10   uint     = 10

	str1   string     = "1"
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", num1, num1)
	fmt.Printf("Type: %T Value: %v\n", num2, num2)
	fmt.Printf("Type: %T Value: %v\n", str1, str1)
	fmt.Printf("Type: %T Value: %v\n", num10, num10)
	
}

//型一覧

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // uint8 の別名

// rune // int32 の別名
//      // Unicode のコードポイントを表す

// float32 float64

// 複素数型
// complex64 complex128
// ほとんど使ってないと思う
//https://www.reddit.com/r/golang/comments/2x7qqo/what_do_you_use_complex64complex128_for/


