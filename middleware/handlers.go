package middleware


import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	// "reflect"
	"github.com/imirjar/notimail_service/models"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()




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
func SendNotimails(w http.ResponseWriter, r *http.Request) {
	//read request
	var notiMail models.NotiMail
	err := json.NewDecoder(r.Body).Decode(&notiMail)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	//to-do	

	//connect to redis db
	ctx := context.Background()
	
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})


	//push mails to redis
	if len(notiMail.Mails.To) > 0 {
		redis_task, _ := json.Marshal(notiMail.Mails)

		err := rdb.LPush(ctx, "notimail:mails", redis_task).Err()		
		if err != nil {
			panic(err)
		}
	}

	//push notifications to redis
	if len(notiMail.Notifications.To) > 0 {
		redis_task, _ := json.Marshal(notiMail.Notifications)

		err := rdb.LPush(ctx, "notimail:notifications", redis_task).Err()
		if err != nil {
			panic(err)
		}
	}
	

	//send response
	res := models.Response{
		Message: "Notification created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

