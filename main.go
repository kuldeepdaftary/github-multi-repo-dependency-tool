package main

import (
    "log"
    "net/http"
    "time"
    "fmt"
)

func main() {
    fmt.Println("Aloha")

    router := NewRouter()

    server := &http.Server{
    	Addr: ":8080",
    	Handler: router,
    	ReadTimeout: 10 * time.Second,
    	WriteTimeout: 10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }

    setupDatabase()

    log.Fatal(server.ListenAndServe())
}
