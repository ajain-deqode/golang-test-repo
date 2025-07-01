package main

import (
    "log"
    "net/http"
    "go-sample-app/internal/handler"
)

func main() {
    http.HandleFunc("/hello", handler.HelloHandler)
    log.Println("Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
