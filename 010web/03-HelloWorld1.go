package main

import (
    "fmt"
    "net/http"
)

func main() {
    //Registering a Request Handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    //80番ポートで動く
    //http://localhost/
    http.ListenAndServe(":80", nil)
}
