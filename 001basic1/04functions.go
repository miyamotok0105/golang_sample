package main

import "fmt"

//goでは0個以上の引数を取れる
//かっこの後ろに戻り値の型を書く
func add(x int, y int) int {
    return x + y
}


func main() {
    fmt.Println(add(42, 13))
}
