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
	"github.com/joho/godotenv"
   	"os"
   	"net/smtp"
)


var ctx = context.Background()



func MailConsumer() {

	err := godotenv.Load(".env")
   	if err != nil {
    	log.Printf("Error while parsing .env file: %v\n", err)
   	}

   	var rdb = redis.NewClient(&redis.Options{
		Addr:	  os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:		 0,
	})



	var mail models.Mail
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_HOST_USER"), os.Getenv("EMAIL_HOST_PASSWORD"), os.Getenv("EMAIL_HOST")) 
    

	for {
		task, err := rdb.LPop(ctx, "notimail:mails").Result()
		if err != nil {
			log.Printf("%v There are no any mails in query. I need some sleep...", err)
			time.Sleep(360 * time.Second)
		} else {
			//to-do with task
			json.Unmarshal([]byte(task), &mail)
			
			
			err := smtp.SendMail(os.Getenv("EMAIL_HOST")+":"+os.Getenv("EMAIL_PORT"), auth, os.Getenv("EMAIL_HOST_USER"), mail.To, []byte(mail.Message)) 
			if err != nil { 
			    log.Fatal(err) 
			}
			
			fmt.Printf("notimail:mail %v\n", mail)
			 
		}
	}
}

func NotificationConsumer() {
	telegramApiToken := "6287615637:AAF85vS9MDnFbuayB9R8VeTDFjMn8hcxKiQ"
	//var telegramChatId		= "1939907187"
	telegramGroupId	 := "-688697681"

	err := godotenv.Load(".env")
   	if err != nil {
    	log.Printf("Error while parsing .env file: %v\n", err)
   	}

   	var rdb = redis.NewClient(&redis.Options{
		Addr:	  os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:		  0,
	})

	var notification models.Notification

	for {
		task, err := rdb.LPop(ctx, "notimail:notifications").Result()
		if err != nil {
			log.Printf("%v There are no any notifications in query. I need some sleep...", err)
			time.Sleep(360 * time.Second)
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