package middleware


import (
	"encoding/json"
	"net/http"
)


type response struct {
	Message string `json:"message,omitempty"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// format a response object
	res := response{
		Message: "Stock created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}