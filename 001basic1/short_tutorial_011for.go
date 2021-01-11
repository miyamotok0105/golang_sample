
package main

import "fmt"

func main() {

    var x int = 1
    var y int = 3

    for x < y {
        fmt.Println(x)
        x++
    }

    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    n := 0
    for {
        n++
        if n > 10 {
            break
        } else if n % 2 == 1 {
            continue
        } else {
            fmt.Println(n)
        }
    }

    colors := [...]string{"Red", "Green", "Blue"}

    for i, color := range colors {
        fmt.Printf("%d: %s\n", i, color)
    }


}

