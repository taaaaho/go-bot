package usecase

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/taaaaho/go-bot/infrastructure/repository"
)

// AddContentsBot ... add content
func AddContentsBot(bot *linebot.Client, event *linebot.Event) {
	var replyMessage string

	repo := repository.NewShoppingList()
	content := event.Message.(*linebot.TextMessage).Text

	// 既に登録されていれば削除、無ければ登録
	if repo.Delete(content) {
		replyMessage = fmt.Sprintf("削除したよ：%s", content)
	} else {
		if err := repo.Add(content); err != nil {
			fmt.Println(err)
			return
		}
		replyMessage = fmt.Sprintf("登録したよ：%s", content)
	}

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Print(err)
	}
}
