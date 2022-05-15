package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/aasumitro/svc-tgbotgo/src/domain"
	"log"
)

type botHandler struct {
	telegramBot *gotgbot.Bot
	userRepository domain.UserRepositoryContract
}

func NewBotHandler(
	bot *gotgbot.Bot,
	userRepo domain.UserRepositoryContract,
) *botHandler {
	return &botHandler{
		telegramBot: bot,
		userRepository: userRepo,
	}
}

func (handler botHandler) Watch()  {
	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher
	handler.dispatch(dispatcher)

	// Start receiving updates.
	if err := updater.StartPolling(
		handler.telegramBot,
		&ext.PollingOpts{DropPendingUpdates: true},
	); err != nil {
		log.Println("Telegram bot : " + err.Error())
	}

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}

func (handler botHandler) dispatch(dispatcher *ext.Dispatcher)  {
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("profile", handler.profile))
	dispatcher.AddHandler(handlers.NewMessage(filters.All, handler.echo))
}