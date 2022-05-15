package handlers

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func (handler botHandler) profile(bot *gotgbot.Bot, ctx *ext.Context) error {
	var msg = "Akun Telegram dan Sistem OkSetda Absensi **[TELAH]** terintegrasi"
	if _, err := handler.userRepository.Profile(
		fmt.Sprintf("%v", ctx.EffectiveChat.Id),
	); err != nil {
		msg = "Akun Telegram dan Sistem OkSetda Absensi **[BELUM]** terintegrasi"
	}

	if _, err := bot.SendMessage(ctx.EffectiveChat.Id, msg, nil); err != nil {
		log.Println("Telegram bot : " + err.Error())
	}

	return nil
}

