package events

import (
	"fmt"
	v1 "github.com/antony0016/introduction-bot/events/v1"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetMessageByEvent(bot *linebot.Client, event *linebot.Event) *linebot.TemplateMessage {
	var message *linebot.TemplateMessage
	var template linebot.Template
	var altText string
	switch event.Type {
	case linebot.EventTypeMessage:
		fmt.Println("received message")
		introductionText := "看看下面的選項或是直接點進我的完整版履歷瞭解我吧!"
		altText = "歡迎來到我的自我介紹!"
		template = v1.IndexTemplate("https://i.imgur.com/MrnL75e.jpg", "歡迎來到我的自我介紹", introductionText)
		message = linebot.NewTemplateMessage(altText, template)
		//fmt.Println(event.Source)

	case linebot.EventTypePostback:
		fmt.Println("received postback")
		command := event.Postback.Data
		if command[0] == '@' {
			altText, template = v1.GetTemplateByOption(command[1:])
		} else if command[0] == '%' {
			altText, template = v1.GetTemplateByScore(command[1:])
		} else if command[0] == '*' {
			altText, template = v1.GetTemplateByScore(command)
			v1.SendScoreToMyGroup(bot, command[1:])
		}
		message = linebot.NewTemplateMessage(altText, template)

	case linebot.EventTypeFollow:
		fmt.Println("following message")
		introductionText := "歡迎來到我的自我介紹，可以點擊下面的各個選項，選擇你想看的部分喔!"
		altText = "歡迎來到我的自我介紹!"
		template = v1.IndexTemplate("https://i.imgur.com/MrnL75e.jpg", "歡迎來到我的自我介紹", introductionText)
		message = linebot.NewTemplateMessage(altText, template)
	}
	return message
}
