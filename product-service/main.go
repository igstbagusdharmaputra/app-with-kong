package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", ProductHandler)

    fmt.Println("Server started at port 8085")
    log.Fatal(http.ListenAndServe(":8085", nil))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Hello, Product Service\n")
}