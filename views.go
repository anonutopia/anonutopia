package main

import (
	"fmt"
	"strings"

	"github.com/anonutopia/gowaves"
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
	msgArr := strings.Fields(tu.Message.Text)
	var msg tgbotapi.Chattable
	send := true

	if len(msgArr) == 1 {
		if msgArr[0] == "/gimme@AnonsRobot" {
			msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), "Adresa novčanika je obavezna. Pokušaj ponovo tako da upišeš i nju (/gimme@AnonsRobot adresa).")
		} else {
			send = false
		}
	} else {
		if msgArr[0] == "/gimme@AnonsRobot" {
			addr := msgArr[1]
			avr, err := wnc.AddressValidate(addr)
			if err == nil {
				if avr.Valid {
					user := &User{Address: addr}
					db.First(user, user)

					if user.ReceivedFreeAnote {
						msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), "Tvoja besplatna Anota već je aktivirana. Morat ćeš unaprijediti svoje hakerske vještine. 😆")
					} else {
						atr := &gowaves.AssetsTransferRequest{
							Amount:    100000000,
							AssetID:   "4zbprK67hsa732oSGLB6HzE8Yfdj3BcTcehCeTA1G5Lf",
							Fee:       100000,
							Recipient: addr,
							Sender:    conf.NodeAddress,
						}

						_, err := wnc.AssetsTransfer(atr)
						if err != nil {
							msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), fmt.Sprintf("Dogodio se problem: %s", err))
						} else {
							user.ReceivedFreeAnote = true
							db.Save(user)
							msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), "Poslao sam ti tvoju 1 besplatnu Anotu! Anonutopia ti želi dobrodošlicu! 🚀")

							if len(user.Referral) > 0 {
								atr := &gowaves.AssetsTransferRequest{
									Amount:    20000000,
									AssetID:   "4zbprK67hsa732oSGLB6HzE8Yfdj3BcTcehCeTA1G5Lf",
									Fee:       100000,
									Recipient: user.Referral,
									Sender:    conf.NodeAddress,
								}

								wnc.AssetsTransfer(atr)
							}
						}
					}
				} else {
					msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), "Nešto nije u redu s adresom tvog novčanika. Molim te da ju provjeriš.")
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
