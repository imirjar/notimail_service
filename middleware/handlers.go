package middleware


import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"strconv"
	// "reflect"
	"github.com/imirjar/notimail_service/models"
	"github.com/redis/go-redis/v9"
	"github.com/joho/godotenv"
   	"os"
)


var ctx = context.Background()


func GetStatus(w http.ResponseWriter, r *http.Request) {


	err := godotenv.Load(".env")
   	if err != nil {
    	log.Printf("Error while parsing .env file: %v\n", err)
   	}

   	var rdb = redis.NewClient(&redis.Options{
		Addr:	  os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:		 0,
	})

	statusMails, err := rdb.LRange(ctx, "notimail:mails", 0, -1).Result()
	if err != nil {
		fmt.Printf("%v There are no any mails in redis", err)
	}

	statusNotifications, err := rdb.LRange(ctx, "notimail:notifications", 0, -1).Result()
	if err != nil {
		fmt.Printf("%v There are no any notifications in redis", err)
	} 

	res := models.Response{
		Message: "Состояние: Активно \n" +"Писем в очереди: "+strconv.Itoa(len(statusMails)) + "\n" + "Уведомлений в очереди: " + strconv.Itoa(len(statusNotifications)) + "\n",
	}

	json.NewEncoder(w).Encode(res)
}


//curl -d '{"mail":{"to":["user@mail.ru", "client@mail.ru"], "subject":"greeting", "message":"Hello my dear friend!"}, "notification":{"to":["123123", "234234"], "messenger":"vk", "message":"Hello my dear friend!"}}' -H "Content-Type: application/json" -X POST http://localhost:8080/send_notimail

func SendNotimails(w http.ResponseWriter, r *http.Request) {
	//read request

	err := godotenv.Load(".env")
   	if err != nil {
    	log.Printf("Error while parsing .env file: %v\n", err)
   	}

   	var rdb = redis.NewClient(&redis.Options{
		Addr:	  os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:		 0,
	})

	var notiMail models.NotiMail
	err = json.NewDecoder(r.Body).Decode(&notiMail)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	//push mails to redis
	if len(notiMail.Mail.To) > 0 {
		redis_task, _ := json.Marshal(notiMail.Mail)
		err := rdb.LPush(ctx, "notimail:mails", redis_task).Err()		
		if err != nil {
			panic(err)
		}
	}

	//push notifications to redis
	if len(notiMail.Notification.To) > 0 {
		redis_task, _ := json.Marshal(notiMail.Notification)
		err := rdb.LPush(ctx, "notimail:notifications", redis_task).Err()
		if err != nil {
			panic(err)
		}
	}

	//send response
	res := models.Response{
		Message: "Notification/mails added to redis query",
	}

	json.NewEncoder(w).Encode(res)
}

