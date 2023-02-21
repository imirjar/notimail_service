package router

import (
	"github.com/imirjar/notimail_service/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/status", middleware.GetStatus).Methods("GET")
	router.HandleFunc("/send_notimail", middleware.SendNotimails).Methods("POST")


	return router
}