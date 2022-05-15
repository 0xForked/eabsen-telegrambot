package config

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"log"
	"net/http"
)

var telegramBot *gotgbot.Bot

func (config AppConfig) InitTelegramBot() {
	log.Println("Initialize telegram bot . . .")
	bot, err := gotgbot.NewBot(config.GetTelegramKey(), &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})

	if err != nil {
		log.Panicln(fmt.Sprintf("telegram bot: %s", err.Error()))
	}

	log.Println("telegram bot: connected")
	setTelegramBot(bot)
}

func setTelegramBot(bot *gotgbot.Bot) {
	telegramBot = bot
}

func (config AppConfig) GetTelegramBot() *gotgbot.Bot {
	return telegramBot
}