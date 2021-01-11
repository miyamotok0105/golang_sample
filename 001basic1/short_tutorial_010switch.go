
package main

import "fmt"

func main() {

    var x int = 1
    var y int = 3
    var mode string = "running"
    var dayOfWeek string = "sat"

    switch mode {
    case "running":
        fmt.Printf("実行中 \n")
    case "stop":
        fmt.Printf("停止中 \n")
    default:
        fmt.Printf("不明 \n")
    }

    switch {
    case x > y:
        fmt.Printf("Big \n")
    case x < y:
        fmt.Printf("Small \n")
    default:
        fmt.Printf("Equal \n")
    }

    //break文はいらない
    switch dayOfWeek {
    case "sat":
        fallthrough
    case "sun":
        fmt.Printf("Horiday \n")
    default:
        fmt.Printf("Weekday \n")
    }

}

