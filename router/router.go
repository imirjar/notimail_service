package router

import (
	"github.com/imirjar/notimail_service/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", middleware.GetStatus).Methods("GET", "OPTIONS")
	router.HandleFunc("/send_mails", middleware.SendMails).Methods("POST", "OPTIONS")
	router.HandleFunc("/send_notifications", middleware.SendNotifications).Methods("POST", "OPTIONS")
	//router.HandleFunc("/send_notifications", middleware.SendNotifications).Methods("POST", "OPTIONS")


	return router
}