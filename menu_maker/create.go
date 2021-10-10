package main

import (
	"fmt"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// createRichMenu()
	listRichMenu()
}

func checkDefaultMenu() {
	bot := createBot()
	res, err := bot.GetDefaultRichMenu().Do()
	if res == nil {
		fmt.Println("err NILだ")
		return
	}
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	println(res.RichMenuID)
}
func setDefaultMenu() {
	bot := createBot()
	res, err := bot.SetDefaultRichMenu("richmenu-72097beaef3bd2dc889cd2cb3f94a9cc").Do()
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	println(res.RequestID)
}
func createBot() *linebot.Client {
	fmt.Println("Create Bot")
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	return bot
}

func twoColumnMenu() linebot.RichMenu {
	fmt.Println("Create Two Column Menu")
	richmenu := linebot.RichMenu{
		Size: linebot.RichMenuSize{
			Width:  2500,
			Height: 843,
		},
		Selected:    true,
		Name:        "Cat Menu",
		ChatBarText: "Tap to nya-",
		Areas: []linebot.AreaDetail{
			linebot.AreaDetail{
				linebot.RichMenuBounds{
					X:      0,
					Y:      0,
					Width:  1250,
					Height: 843,
				},
				linebot.RichMenuAction{
					Type: "postback",
					Data: "action=list",
				},
			},
			linebot.AreaDetail{
				linebot.RichMenuBounds{
					X:      843,
					Y:      0,
					Width:  1250,
					Height: 843,
				},
				linebot.RichMenuAction{
					Type: "postback",
					Data: "action=clear",
				},
			},
		},
	}
	return richmenu
}
func listRichMenu() {
	fmt.Println("List RichMenu")
	bot := createBot()
	res, err := bot.GetRichMenuList().Do()
	if err != nil {
		fmt.Errorf("err %v", err)
	}

	for _, richMenu := range res {
		println(richMenu.RichMenuID)
	}
}

func createRichMenu() {
	fmt.Println("Create Menu")
	richmenu := twoColumnMenu()
	bot := createBot()
	res, err := bot.CreateRichMenu(richmenu).Do()
	if err != nil {
		fmt.Errorf("err %v", err)
	}

	println(res.RichMenuID)
}
