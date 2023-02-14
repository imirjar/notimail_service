package observer

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/imirjar/notimail_service/models"
	"context"
	"time"
	"encoding/json"
	"log"
	"net/http"
)

var mail 			models.Mail
var notification 	models.Notification

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

var telegramApiToken 	= "6287615637:AAF85vS9MDnFbuayB9R8VeTDFjMn8hcxKiQ"
//var telegramChatId		= "1939907187"
var telegramGroupId		= "-688697681"


func MailConsumer() {
	
	for {
		task, err := rdb.LPop(ctx, "notimail:mails").Result()
		if err != nil {
			log.Printf("%v There are no any tasks in query. I need some sleep...", err)
			time.Sleep(6000 * time.Millisecond)
		} else {
			//to-do with task
			json.Unmarshal([]byte(task), &mail)



			fmt.Printf("notimail:mails %v\n", mail)
		}
		
	}
}

func NotificationConsumer() {
	
	for {
		task, err := rdb.LPop(ctx, "notimail:notifications").Result()
		if err != nil {
			log.Printf("%v There are no any tasks in query. I need some sleep...", err)
			time.Sleep(6000 * time.Millisecond)
		} else {
			//to-do with task
			json.Unmarshal([]byte(task), &notification)
			fmt.Printf("notimail:notification %v\n", notification)

			messengerApiUrl := "https://api.telegram.org/bot"+telegramApiToken+"/sendMessage?chat_id="+telegramGroupId+"&text="+notification.Message

			resp, err := http.Get(messengerApiUrl)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(resp)
			//https://api.telegram.org/bot6287615637:AAF85vS9MDnFbuayB9R8VeTDFjMn8hcxKiQ/sendMessage?chat_id=1939907187&text='%v'
		}
	}

}