package usecase

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/taaaaho/go-bot/infrastructure/repository"
)

type goods struct {
	Goods string `firestore:"goods"`
}

// GetShoppingList ... get list
func GetShoppingList(bot *linebot.Client, event *linebot.Event) {
	var replayMessage string

	repo := repository.NewShoppingList()
	shoppingList, err := repo.List()
	if err != nil {
		fmt.Printf("Error when get shopping list %v", err)
		replayMessage = "もう一度試してみて"
	}

	for _, v := range shoppingList {
		if replayMessage != "" {
			replayMessage = replayMessage + "\n・" + v.Contents
		} else {
			replayMessage = fmt.Sprintf("ねここのお使いリスト\n・%s", v.Contents)
		}
	}

	if replayMessage == "" {
		replayMessage = "お買い物リストが無いよ"
	}

	if _, err := bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(replayMessage),
	).Do(); err != nil {
		log.Print(err)
	}
}
