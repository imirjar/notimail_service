package main

import (
	"github.com/gorilla/mux"
	"github.com/imirjar/notimail_service/middleware"
    "net/http"
)

// Router is exported and used in main.go
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    http.Handle("/", r)
}