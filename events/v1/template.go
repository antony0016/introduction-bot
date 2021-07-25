package v1

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// IndexTemplate : return main menu of introduction, projects etc.
func IndexTemplate(thumbnailImageURL, title, text string) linebot.Template {
	var template linebot.Template
	// text in PostbackAction already deprecated, just leave it empty.

	// sampleAction := linebot.NewPostbackAction("sample", "@sample", "", "give me sample")
	resume := linebot.NewPostbackAction("完整版履歷", "@resume", "", "")
	skills := linebot.NewPostbackAction("我的技能", "@skills", "", "")
	projects := linebot.NewPostbackAction("作品集及其他資料", "@projects", "", "")
	//others := linebot.NewPostbackAction("其他", "@others", "", "")
	scores := linebot.NewPostbackAction("幫我評分", "%scores", "", "")
	template = linebot.NewButtonsTemplate(thumbnailImageURL, title, text,
		resume, skills, projects, scores,
	)

	return template
}
