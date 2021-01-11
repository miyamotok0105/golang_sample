
package main

import "fmt"

func main() {

    var x int = 1
    var y int = 3
    var max int = 2

    if x > max {
        max = x
    }

    if x > y {
        fmt.Printf("Big \n")
    } else if x < y {
        fmt.Printf("Small \n")
    } else {
        fmt.Printf("Equal \n")
    }
    fmt.Printf(string(max)+"\n")
}

