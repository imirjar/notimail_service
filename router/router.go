package router

import (
	"github.com/imirjar/notimail_service/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", middleware.GetStatus).Methods("GET", "OPTIONS")
	//router.HandleFunc("/send_mail", middleware.GetStatus).Methods("POST", "OPTIONS")


	return router
}