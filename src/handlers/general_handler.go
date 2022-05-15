package handlers

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/aasumitro/svc-tgbotgo/src/helpers"
	"log"
)

func start(bot *gotgbot.Bot, ctx *ext.Context) error {
	if _, err := bot.SendMessage(ctx.EffectiveChat.Id, fmt.Sprintf(
		helpers.CmdStartMsg,
		ctx.EffectiveChat.Username,
		bot.User.Username,
		ctx.EffectiveChat.Id,
	), nil); err != nil {
		log.Println("Telegram bot : " + err.Error())
	}

	return nil
}

func (handler botHandler) echo(bot *gotgbot.Bot, ctx *ext.Context) error {
	if len(ctx.EffectiveMessage.Text) >= 7 && ctx.EffectiveMessage.Text[0:7] == "CONNECT" {
		data := helpers.Explode("#", ctx.EffectiveMessage.Text)
		msg := "Integrasi Telegram dan OkSetda Absensi **[BERHASIIL]**"
		if _, err := handler.userRepository.Connect(
			fmt.Sprintf("%v", data[1]),
			fmt.Sprintf("%v", data[2]),
			fmt.Sprintf("%v", ctx.EffectiveChat.Id),
		); err != nil {
			 msg = fmt.Sprintf(
			 	"%s",
			 	err)
		}

		if _, err := bot.SendMessage(ctx.EffectiveChat.Id, msg, nil); err != nil {
			log.Println("Telegram bot : " + err.Error())
		}
	} else {
		if _, err := ctx.EffectiveMessage.Reply(
			bot,
			fmt.Sprintf(
				helpers.CmdNotFound,
				bot.User.Username,
				ctx.EffectiveMessage.Text,
			),
			nil); err != nil {
			log.Println("Telegram bot : " + err.Error())
		}
	}

	return nil
}
