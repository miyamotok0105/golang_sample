
package main

import "fmt"

func main() {
    // スライス作成
    a1 := []string{}			// スライス。個数不定
    a1 = append(a1, "Red")
    a1 = append(a1, "Green")
    a1 = append(a1, "Blue")
    fmt.Println(a1[0], a1[1], a1[2])

    //スライスはmakeを用いたメモリ確保ができる
    bufa := make([]byte, 0, 1024)
    fmt.Println(bufa)
}

