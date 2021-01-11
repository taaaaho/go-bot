package usecase

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/taaaaho/go-bot/infrastructure/repository"
)

// GetShoppingList ... get list
func GetPoint(bot *linebot.Client, event *linebot.Event) {
	var replayMessage string

	repo := repository.NewShoppingList()
	point := repo.GetPoint(event.Source.UserID)

	replayMessage = fmt.Sprintf("あなたの保有ポイント：%d", point)

	if _, err := bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(replayMessage),
	).Do(); err != nil {
		log.Print(err)
	}
}
