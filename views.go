package main

import (
	"fmt"
	"strings"

	"gopkg.in/macaron.v1"
	"gopkg.in/telegram-bot-api.v4"
)

func homeView(ctx *macaron.Context) {
	ctx.Data["Title"] = ""

	ctx.HTML(200, "home")
}

func profitView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Making Profit in Anonutopia | "

	ctx.HTML(200, "profit")
}

func anoteView(ctx *macaron.Context) {
	ctx.Data["Title"] = "What is ANOTE? | "

	ctx.HTML(200, "anote")
}

func projectsView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Anonutopia Projects | "

	ctx.HTML(200, "projects")
}

func contactView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Contact | "

	ctx.HTML(200, "contact")
}

func webhookView(ctx *macaron.Context, tu TelegramUpdate) {
	msgArr := strings.Split(tu.Message.Text, " ")
	var msg tgbotapi.Chattable
	send := true

	if len(msgArr) == 1 {
		if msgArr[0] == "/gimme@AnonsRobot" {
			msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), "Niste upisali adresu svog novčanika. Pokušajte ponovo tako da upišete i nju (/gimme@AnonsRobot adresa).")
		} else {
			send = false
		}
	} else {
		if msgArr[0] == "/gimme@AnonsRobot" {
			addr := msgArr[1]
			avr, err := wnc.AddressValidate(addr)
			if err == nil {
				if avr.Valid {
					msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), "Poslao sam vam vašu 1 besplatnu Anotu! Dobrodošli u Anonutopiju! 🚀")
				} else {
					msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), "Nešto nije u redu s adresom vašeg novčanika. Molim vas da ju provjerite.")
				}
			} else {
				msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), fmt.Sprintf("Dogodio se problem: %s", err))
			}
		} else {
			send = false
		}
	}

	if send {
		bot.Send(msg)
	}
}
