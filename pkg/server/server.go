package server

import (
    "fmt"
    "log"
    "net/http"
)

func StartServer() {
    fmt.Println("Starting server at http://localhost:3333")
    http.Handle("/", http.FileServer(http.Dir("./public")))
    log.Fatal(http.ListenAndServe(":3333", nil))
}

