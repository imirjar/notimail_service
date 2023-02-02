package main

import (
	"github.com/imirjar/notimail_service/router"
    "fmt"
    "log"
    "net/http"
)

// Router is exported and used in main.go
func main() {
   r := router.Router()
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)
    fmt.Println("Starting server on the port 8080...")

    log.Fatal(http.ListenAndServe(":8080", r))
}