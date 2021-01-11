package usecase

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/taaaaho/go-bot/infrastructure/repository"
)

// DeleteShoppingList ... reset list
func DeleteShoppingList(bot *linebot.Client, event *linebot.Event) {
	repo := repository.NewShoppingList()
	if _, err := repo.DeleteAll(); err != nil {
		fmt.Println(err)
	}

	var replayMessage = "お買い物リストを削除したよ"
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replayMessage)).Do(); err != nil {
		log.Print(err)
	}
}
