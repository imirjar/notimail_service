package middleware


import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"github.com/imirjar/notimail_service/models"
)


func GetStatus(w http.ResponseWriter, r *http.Request) {
	// format a response object
	fmt.Printf("%v\n", r)
	res := models.Response{
		Message: "Stock created successfully",
	}
	// send the response
	json.NewEncoder(w).Encode(res)
}

//curl -d '{"to":["user@mail.ru", "client@mail.ru"], "subject":"greeting", "message":"Hello my dear friend!"}' -H "Content-Type: application/json" -X POST http://localhost:8080/send_mails
func SendMails(w http.ResponseWriter, r *http.Request) {

	//read request
	var mail models.Mail//создаем переменную с типом данных "структура письма"
	err := json.NewDecoder(r.Body).Decode(&mail)//декодируем двоичные данные в участок памяти переменной mail

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	//to-do
	for i := 0; i < len(mail.To); i++ {
		fmt.Printf("%v\n", mail.To[i])
	}

	//send response
	res := models.Response{
		Message: "Mail created successfully",
	}
	json.NewEncoder(w).Encode(res)
}


//curl -d '{"to":["123", "345"], "message":"Hello my dear friend!"}' -H "Content-Type: application/json" -X POST http://localhost:8080/send_notifications
func SendNotifications(w http.ResponseWriter, r *http.Request) {

	//read request
	var notification models.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	//to-do	
	for i := 0; i < len(notification.To); i++ {
		fmt.Printf("%v\n", notification.To[i])
	}

	//send response
	res := models.Response{
		Message: "Notification created successfully",
	}
	json.NewEncoder(w).Encode(res)
}





