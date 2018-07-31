package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func initBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(conf.Telegram)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(-304575934, "Robot successfully started.")
	bot.Send(msg)

	return bot
}
