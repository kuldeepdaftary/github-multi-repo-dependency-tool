package main

import (
    "fmt"
    "net/http"
)

func serverTest(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Server is up and running")
}
