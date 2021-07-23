package events

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetMessageByEvent(event *linebot.Event) *linebot.TemplateMessage {
	var message *linebot.TemplateMessage
	switch event.Type {
	case linebot.EventTypeMessage:
		fmt.Println("received message")
	case linebot.EventTypePostback:
		fmt.Println("received postback")
	case linebot.EventTypeFollow:
		fmt.Println("following message")
	}
	return message
}
