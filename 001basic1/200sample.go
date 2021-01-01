package main

import (
		"fmt"
		"time"
)

//コンポジット型は複数のデータを１つの集合として表す型。
//構造体 struct,配列 array,スライス slice、マップ map,チャネル channel, chanがある。
func run1() {
		//フィールドを持たない構造体
		var empty struct{}

		//フィールド３つを持つ構造体
		var point struct {
			ID string
			x, y int
		}
		fmt.Println(empty)
		fmt.Println(point)
}

func run2() {
		//配列の定義
		//ゼロで初期化
		var array [5]int
		//５つの要素を持つ配列を定義
		arrayLiteral := [5]int{1,2,3,4,5}
		//要素数から配列数を推論。この場合は５
		arrayInterface := [...]int{1,2,3,4,5}
		//配列のインデックス と値を指定
		//インデックス の指定がない場合はゼロ
		arrayIndex := [...]int{2:1, 5:5, 7:13}
		fmt.Println(array)
		fmt.Println(arrayLiteral)
		fmt.Println(arrayInterface)
		fmt.Println(arrayIndex)
		//配列は要素数を変えれなくて、スライスは変えれて型が違う場合があるやつ
}

func run3() {
		//ゼロ値で初期化
		var slice []int
		//５つの要素を持つスライスを定義
		sliceLiteral := []int{1,2,3,4,5}
		//ゼロ値で初期化
		var m map[string]int
		//２つの要素を持つマップを定義
		mapLiteral := map[string]int {
				"John": 42,
				"Tom":33,
		}
		fmt.Println(slice)
		fmt.Println(sliceLiteral)
		fmt.Println(m)
		fmt.Println(mapLiteral)
		//スライスはコンポジット型でもよく使う型
}

func run4() {
		//新しい型を定義
		type MyDuration time.Duration
		d := MyDuration(100)

		fmt.Println("%T", d)

		td := time.Duration(d)

		md := 100 * d

		fmt.Println("td: %T, md: %T", td, md)
}

func run5() {
		src := []int{1,2,3,4,5}
		// lenとcapでスライスの長さと容量を表示
		fmt.Println("lenとcapでスライスの長さと容量を表示")
		fmt.Println(src, len(src), cap(src))

		src = append(src, 5)
		fmt.Println(src, len(src), cap(src))

		//スライスを初期化
		fmt.Println("スライスを初期化")
		sliceMake := make([]int, 2, 3)
		sliceIndex := []int{2:1, 5:5, 7:13}
		fmt.Println(sliceMake)
		fmt.Println(sliceIndex)

		//copyでスライスを複製
		fmt.Println("copyでスライスを複製")
		dst := make([]int, len(src))
		copy(dst, src)
		fmt.Println(dst, len(dst), cap(dst))


}

func run6() {
		//appendでスライス同士を連結
		fmt.Println("appendでスライス同士を連結")
		src1, src2 := []int{1, 2}, []int{3,4,5}
		dst := append(src1, src2...)
		fmt.Println(dst, len(dst), cap(dst))

		//append,copyでスライス内の要素を削除
		// fmt.Println("append,copyでスライス内の要素を削除")
		// i := 2
		// dst := append(src[:i], src[i+1:]...)
}

func main() {
		fmt.Println("run1=====")
		run1()
		fmt.Println("run2=====")
		run2()
		fmt.Println("run3=====")
		run3()
		fmt.Println("run4=====")
		run4()
		fmt.Println("run5=====")
		run5()
		fmt.Println("run6=====")
		run6()
}

