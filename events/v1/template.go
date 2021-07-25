package v1

import "github.com/line/line-bot-sdk-go/v7/linebot"

func indexActions() []linebot.PostbackAction {
	var a []linebot.PostbackAction
	introduction := *linebot.NewPostbackAction("自我介紹label", "@introduction", "自我介紹text", "自我介紹display text")
	projects := *linebot.NewPostbackAction("作品集label", "@introduction", "作品集text", "作品集display text")
	a = append(a, introduction)
	a = append(a, projects)

	return a
}
