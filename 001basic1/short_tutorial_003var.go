
package main

import "fmt"

func main() {
    // var a1 int
    // var a1 int = 123
    // var a2 = 123
    // a3 := 123
    var (
        a1 int = 123
        // a2 int = 456
        name string
        age int
    )
    a1 = 456
    name, age = "Yamada", 26
    fmt.Printf("num=%d str=%s\n", a1, name)
    fmt.Printf("num=%d str=%s\n", age, name)
}

