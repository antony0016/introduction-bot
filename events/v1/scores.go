package v1

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"os"
)

func SendScoreToMyGroup(bot *linebot.Client, score string) {
	result := fmt.Sprintf("獲得評分: %s", score)
	fmt.Println(os.Getenv("GROUP_ID"))
	if _, err := bot.PushMessage(os.Getenv("GROUP_ID"), linebot.NewTextMessage(result)).Do(); err != nil {
		log.Println(err)
	}
}
