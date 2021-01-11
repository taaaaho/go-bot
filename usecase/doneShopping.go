package usecase

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/taaaaho/go-bot/infrastructure/repository"
)

// DoneShopping ... bulk send
func DoneShopping(bot *linebot.Client, event *linebot.Event) {
	repo := repository.NewShoppingList()
	count, err := repo.DeleteAll()
	if err != nil {
		fmt.Println(err)
	}

	err = repo.AddPoint(event.Source.UserID, count)
	if err != nil {
		log.Print(err)
	}

	replayMessage := fmt.Sprintf("お使いしてきたよ\n・お使い点数：%d", count)
	if _, err := bot.BroadcastMessage(linebot.NewTextMessage(replayMessage)).Do(); err != nil {
		log.Print(err)
	}

	replayMessage = fmt.Sprintf("ポイントゲット：%d", count)
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replayMessage)).Do(); err != nil {
		log.Print(err)
	}
}
