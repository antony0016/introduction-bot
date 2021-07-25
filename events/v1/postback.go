package v1

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetTemplateByOption(command string) (string, linebot.Template) {
	var template linebot.Template
	var altText string = "..."
	homeAction := linebot.NewPostbackAction("選單", "@menu", "", "")
	fmt.Println("command: " + command)
	switch {
	case command == "resume":
		altText = "我的完整版履歷"
		introduction := "點擊下面的選項，就可以看到完整版的履歷喔!"
		template = linebot.NewButtonsTemplate("https://i.imgur.com/DRv3YdJm.jpg", "自我介紹", introduction,
			linebot.NewURIAction(
				"點我前往👈",
				"https://www.canva.com/design/DAEjXuq6ldc/0P3CjXn7qfdB2rSB3bGYnA/view?utm_content=DAEjXuq6ldc&utm_campaign=designshare&utm_medium=link&utm_source=publishsharelink#1",
			),
		)
	case command == "skills":
		altText = "我的技能樹!"
		frontend := linebot.NewCarouselColumn("https://i.imgur.com/1KtDKVGm.png",
			"前端技能", "語言:HTML, CSS, JavaScript \n框架:VueJS",
			homeAction,
		)
		backend := linebot.NewCarouselColumn("https://i.imgur.com/4vRIlcsm.png",
			"後端技能", "語言:Golang, Python \n框架:Gin, Flask, ",
			homeAction,
		)
		others := linebot.NewCarouselColumn("https://i.imgur.com/UloljLYm.png",
			"其他技能", "語言:C++, Java\n框架:Selenium\n服務:, Docker, Git, LineBot",
			homeAction,
		)
		template = linebot.NewCarouselTemplate(frontend, backend, others)
	case command == "projects":
		altText = "我的作品集與其他有利審查資料"
		swSystem := linebot.NewCarouselColumn("https://i.imgur.com/XPcY0Yim.png",
			"機車行客戶管理系統", "整合客戶資料以及工單的管理系統。使用技術: VueJS, Golang, Docker",
			homeAction,
		)
		swSystemDemo := linebot.NewCarouselColumn("https://i.imgur.com/XPcY0Yim.png",
			"機車行客戶管理系統", "點擊下方\"打開Demo\"可以連結到Demo網站試用",
			linebot.NewURIAction("打開Demo", "https://demo.matsuno-seki.com"),
		)
		line := linebot.NewCarouselColumn("https://i.imgur.com/Mx9UIys.jpg",
			"自我介紹 line bot", "點擊下方\"前往 Github\" 可以在 Github 查看完整程式碼",
			linebot.NewURIAction("前往 Github", "https://github.com/antony0016/introduction-bot"),
		)
		aacsb := linebot.NewCarouselColumn("https://i.imgur.com/455BMyYm.png",
			"AACSB教授評分系統", "讀取 excel 教授資料，以程式自動化評分並寫入資料庫，大幅減少人力及時間支出。",
			homeAction,
		)
		amazon := linebot.NewCarouselColumn("https://i.imgur.com/YALhjPem.png",
			"商城關鍵字爬蟲(接案)", "於淘寶及Amazon等大型網路商城，篩選同關鍵字內特定商品供案主分析使用。",
			homeAction,
		)
		blog := linebot.NewCarouselColumn("https://i.imgur.com/fB9K4ks.png",
			"我的技術筆記部落格", "將學過的技能做筆記並分享，也用來記錄自己的成長的部落格",
			linebot.NewURIAction("我的部落格", "https://blog.matsuno-seki.com"),
		)
		medium := linebot.NewCarouselColumn("https://i.imgur.com/KRbgJah.png",
			"我的Medium", "基本上和blog的內容是同步的，可根據讀者愛好自行選擇",
			linebot.NewURIAction("我的Medium", "https://feather881030.medium.com/"),
		)
		jlpt := linebot.NewCarouselColumn("https://i.imgur.com/8ZwRvHI.png",
			"JLPT N2", "我很喜歡動漫及日本文化，花了很多時間看動漫，所以也花了點時間準備JLPT並拿到證照了!",
			linebot.NewURIAction("我要看證照!", "https://i.imgur.com/A3yWLMI.jpg"),
		)
		template = linebot.NewCarouselTemplate(swSystem, aacsb, line, amazon, swSystemDemo, blog, medium, jlpt)
	case command == "menu":
		altText = "選單"
		introductionText := "歡迎來到我的自我介紹，可以點擊下面的各個選項，選擇你想看的部分喔!"
		template = IndexTemplate("https://i.imgur.com/MrnL75e.jpg", "選單", introductionText)
	case command == "scores":
		altText = "給我一個分數吧!"
		template = linebot.NewCarouselTemplate()
	}
	if altText == "..." {
		fmt.Println("warning: this command \"", command, "\" didn't change altText")
	}
	return altText, template
}

func GetTemplateByScore(command string) (string, linebot.Template) {
	var template linebot.Template
	var altText string = "..."
	fmt.Println("command: " + command)

	switch {
	case command == "scores":
		altText = "給我一個分數吧😁"
		five := linebot.NewCarouselColumn("https://i.imgur.com/JsDuREIm.jpg", "五顆星", "讚啦，快進來一起創造更多便利的服務!",
			linebot.NewPostbackAction("⭐⭐⭐⭐⭐", "*5", "", ""))
		four := linebot.NewCarouselColumn("https://i.imgur.com/qyo9X5Lm.jpg", "四顆星", "你很不錯喔，很期待你可以進來!",
			linebot.NewPostbackAction("⭐⭐⭐⭐", "*4", "", ""))
		three := linebot.NewCarouselColumn("https://i.imgur.com/ykJGOxTm.jpg", "三顆星", "你還行啦，進來可以學到更多，進來好好努力學習!",
			linebot.NewPostbackAction("⭐⭐⭐", "*3", "", ""))
		two := linebot.NewCarouselColumn("https://i.imgur.com/5qeZeaJm.jpg", "二顆星", "不太行喔，進來好好磨練一下!",
			linebot.NewPostbackAction("⭐⭐", "*2", "", ""))
		one := linebot.NewCarouselColumn("https://i.imgur.com/KmVBJnTm.jpg", "一顆星", "抱歉...",
			linebot.NewPostbackAction("⭐", "*1", "", ""))
		template = linebot.NewCarouselTemplate(five, four, three, two, one)
	default:
		altText = "感謝你的評分!"
		scoresText := "感謝你的評分!\n如果有還要重新看一遍的可以再看一次喔!"
		template = IndexTemplate("https://i.imgur.com/MrnL75e.jpg", "評分完成!", scoresText)
	}

	if altText == "..." {
		fmt.Println("warning: this command \"", command, "\" didn't change altText")
	}
	return altText, template
}
