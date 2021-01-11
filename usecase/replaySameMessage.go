package usecase

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ReplySameMessage ... replay same
func ReplySameMessage(bot *linebot.Client, event *linebot.Event, message string) {
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Print(err)
	}
}
