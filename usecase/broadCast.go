package usecase

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/taaaaho/go-bot/infrastructure/repository"
)

// BroadCast ... bulk send
func BroadCast(bot *linebot.Client, event *linebot.Event) {
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
			replayMessage = fmt.Sprintf("お使いいってきてね\n・%s", v.Contents)
		}
	}

	if _, err := bot.BroadcastMessage(linebot.NewTextMessage(replayMessage)).Do(); err != nil {
		log.Print(err)
	}
}
