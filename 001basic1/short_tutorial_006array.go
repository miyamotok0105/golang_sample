
package main

import "fmt"

func main() {
    // 配列作成
    // var a1 := [3] string{}

    var a1[3] string = [3]string {}
    a1[0] = "Red"
    a1[1] = "Green"
    a1[2] = "Blue"
    fmt.Println(a1[0], a1[1], a1[2])

    // a1 := [3]string{"Red", "Green", "Blue"}
    // fmt.Println(a1[0], a1[1], a1[2])
}

