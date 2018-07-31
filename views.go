package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/anonutopia/gowaves"
	"github.com/go-macaron/i18n"
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
	var lang string

	if tu.Message.Chat.ID == -304575934 || tu.Message.Chat.ID == -1001161265502 || tu.Message.Chat.ID == -1001397587839 {
		lang = "hr"
	} else if tu.Message.Chat.ID == -1001249635625 {
		lang = "sr"
	} else {
		lang = "en-US"
	}

	log.Println(tu.Message.Chat.ID)

	m.Use(i18n.I18n(i18n.Options{
		Langs: []string{lang},
		Names: []string{lang},
	}))

	if len(msgArr) == 1 {
		if msgArr[0] == "/gimme@AnonsRobot" {
			msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ctx.Tr("addressRequired"))
		} else {
			send = false
		}
	} else if len(msgArr) > 1 {
		if msgArr[0] == "/gimme@AnonsRobot" {
			addr := msgArr[1]
			avr, err := wnc.AddressValidate(addr)
			if err == nil {
				if avr.Valid {
					user := &User{Address: addr}
					db.FirstOrCreate(user, user)

					if user.ReceivedFreeAnote {
						msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ctx.Tr("alreadyActivated"))
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
							msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), fmt.Sprintf(ctx.Tr("error"), err))
						} else {
							user.ReceivedFreeAnote = true
							db.Save(user)
							msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ctx.Tr("anoteSent"))

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
					msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ctx.Tr("addressNotValid"))
				}
			} else {
				msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), fmt.Sprintf(ctx.Tr("error"), err))
			}
		} else {
			send = false
		}
	} else {
		send = false
	}

	if send {
		bot.Send(msg)
	}
}
