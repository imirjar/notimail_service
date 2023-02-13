package main

import (
	"github.com/imirjar/notimail_service/router"
   "github.com/imirjar/notimail_service/observer"
   "fmt"
   "log"
   "net/http"
)

// Router is exported and used in main.go
func main() {
   r := router.Router()
   go observer.MailConsumer()
   go observer.NotificationConsumer()
   fmt.Println("Starting server on the port 8080...")
   log.Fatal(http.ListenAndServe(":8080", r))
}