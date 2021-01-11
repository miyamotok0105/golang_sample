
package main

import "fmt"

func main() {
    // type UtcTime string		// string型の別名 UtcTime を定義
    // type JstTime string		// string型の別名 JstTime を定義
    type (
        UtcTime string
        JstTime string
    )
    var t1 UtcTime = "00:00:00"
    var t2 JstTime = "09:00:00"
    // t1 = t2				// 型が異なるので代入エラー
    //型変換
    var a1 uint16 = 1234
    var a2 uint32 = uint32(a1)

    fmt.Printf("%s %s\n", t1, t2)
    fmt.Printf("%d %d\n", a1, a2)
}

