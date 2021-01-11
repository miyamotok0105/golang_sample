
package main

import "fmt"

func main() {
    var status string = "running..."
    funcA(status)

}

func GetFileName(string fileName) (int, string) {
    var err int = 200
    fmt.Printf(fileName)
    return err, fileName
}

func ReadFile(string fileName) (int, string) {
    var err int = 200
    fmt.Printf(fileName)
    return err, fileName
}

func funcA(string status) (int, string) {
    fmt.Printf(status)
    err := nil
    var filename string = ""
    data := ""
    err, filename = GetFileName(filename)
    if err != nil {
        goto Done
    }
    err, data = ReadFile(filename)
    if err != nil {
        goto Done
    }
    fmt.Printf(err, data)
Done:
    fmt.Printf(err, data)
}
