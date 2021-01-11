
package main

import "fmt"

func main() {

    // マップを定義する
    a1 := map[string]int{
        "x": 100,
        "y": 200,			// 改行する場合はカンマ必須
    }

    // マップを参照する
    fmt.Println(a1["x"])

    // マップに要素を加える
    a1["z"] = 300

    // マップの要素を削除する
    delete(a1, "z")

    // マップの長さを求める
    fmt.Println(len(a1))

    // マップに要素が存在するかどうかを調べる
    _, ok := a1["z"]
    if ok {
        fmt.Println("Exist")
    } else {
        fmt.Println("Not exist")
    }

    // マップをループ処理する
    for key, value := range a1 {
        fmt.Printf("%s=%d\n", key, value)
    }


}

