package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/taaaaho/go-bot/pkg/linebotfactory"
	"github.com/taaaaho/go-bot/usecase"
)

// Handler ... handler struct
type Handler struct {
	http.Handler
}

const (
	list  string = "list"
	add   string = "add"
	clear string = "clear"
	broad string = "broad"
	done  string = "done"
	point string = "point"
)

// NewHandler ... Craete chi handler
func NewHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/", postMessage)

	return r
}

func postMessage(w http.ResponseWriter, req *http.Request) {
	bot, err := linebotfactory.Newbot()
	if err != nil {
		fmt.Println(fmt.Errorf("Error in create new bot: %v", err))
		w.WriteHeader(500)
		return
	}
	events, err := bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case list:
					usecase.GetShoppingList(bot, event)
				case add:
					usecase.AddContentsBot(bot, event)
				case clear:
					usecase.DeleteShoppingList(bot, event)
				case broad:
					usecase.BroadCast(bot, event)
				case done:
					usecase.DoneShopping(bot, event)
				case point:
					usecase.GetPoint(bot, event)
				default:
					usecase.AddContentsBot(bot, event)
				}
			case *linebot.StickerMessage:
				replyMessage := fmt.Sprintf(
					"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
