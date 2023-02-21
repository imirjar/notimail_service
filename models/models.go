package models


type Mail struct {
	To 				[]string 		`json:"to"`
	Subject 		string 			`json:"subject"`
	Message 		string 			`json:"message"`
}

type Notification struct {
	To 				[]string 		`json:"to"`
	Messenger		string			`json:"messenger"`
	Message 		string 			`json:"message"`
}

type NotiMail struct {
	Notification 	Notification  	`json:"notification"`
	Mail 			Mail 			`json:"mail"`
}

type Response struct {
	Message 		string 			`json:"message,omitempty"`
}