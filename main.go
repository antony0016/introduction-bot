package main

import (
	"fmt"
	myEvent "github.com/antony0016/introduction-bot/events"
	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
	"os"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Println("Starting server...")
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		replyToken := event.ReplyToken
		message := myEvent.GetMessageByEvent(bot, event)
		if _, err = bot.ReplyMessage(replyToken, message).Do(); err != nil {
			log.Println(err)
		}

	}
}
