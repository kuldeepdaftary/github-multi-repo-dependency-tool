package main

import (
    "log"
    "net/http"
    "time"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Aloha")

    router := NewRouter()

    server := &http.Server{
    	Addr: determineListenAddress(),
    	Handler: router,
    	ReadTimeout: 10 * time.Second,
    	WriteTimeout: 10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }

    setupDatabase()

    log.Fatal(server.ListenAndServe())
}

func determineListenAddress() string {
  port := os.Getenv("PORT")
  if port == "" {
    return ":8080"
  }
  return ":" + port
}
