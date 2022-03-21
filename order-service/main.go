package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", OrderHandler)

    fmt.Println("Server started at port 8084")
    log.Fatal(http.ListenAndServe(":8084", nil))
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Hello, Order Service\n")
}