package linebotfactory

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/line/line-bot-sdk-go/linebot"
)

// env config
type env struct {
	Secret      string `envconfig:"CHANNEL_SECRET"`
	AccessToken string `envconfig:"CHANNEL_ACCESS_TOKEN"`
}

// Newbot ... create bot
func Newbot() (*linebot.Client, error) {
	var env env
	if err := envconfig.Process("", &env); err != nil {
		fmt.Println(fmt.Errorf("Error in parse env: %v", err))
	}

	bot, err := linebot.New(
		env.Secret,
		env.AccessToken,
	)

	return bot, err
}
