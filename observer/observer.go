package observer

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/imirjar/notimail_service/models"
	"context"
	"time"
	"encoding/json"
	"log"
)
// Router is exported and used in main.go
// func Consumer(link <-chan string, done chan<- bool) {
// 	for b := range link {
// 		fmt.Println(b)
// 	}
// 	done <- true
// }
var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

var mail models.Mail
var notification models.Notification


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
		}
	}

}