package models

type Response struct {
	Message string `json:"message,omitempty"`
}

type Mail struct {
	To 		[]string 	`json:"to"`
	Subject string 		`json:"subject"`
	Message string 		`json:"message"`
}


type Notification struct {
	To 		[]string 	`json:"to"`
	Message string 		`json:"message"`
}

