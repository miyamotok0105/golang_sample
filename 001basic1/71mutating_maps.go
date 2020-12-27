package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	//要素追加
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	//要素削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	//存在チェック okの場所
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

