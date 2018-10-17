package main

import (
	"fmt"
	"log"
	"strings"

	ui18n "github.com/Unknwon/i18n"
	"github.com/anonutopia/gowaves"
	"gopkg.in/macaron.v1"
	"gopkg.in/telegram-bot-api.v4"
)

func homeView(ctx *macaron.Context) {
	ctx.Data["Title"] = ""

	ctx.HTML(200, "home")
}

func crowdfundingView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Anonutopia Crowdfunding | "

	ctx.HTML(200, "crowdfunding")
}

func aboutView(ctx *macaron.Context) {
	ctx.Data["Title"] = "About Anonutopia | "

	ctx.HTML(200, "about")
}

func applyView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Apply for Citizenship | "

	ctx.HTML(200, "apply")
}

func cryptocountriesView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Crypto Countries | "

	ctx.HTML(200, "cryptocountries")
}

func tagsView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Anonutopia Tags | "

	ctx.HTML(200, "tags")
}

func docsView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Anonutopia Docs | "

	ctx.HTML(200, "docs")
}

func stakingView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Staking (Leasing) with Most Profitable Waves Node | "

	ctx.HTML(200, "staking")
}

func airdropView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Claim 1 Free ANOTE | "

	ctx.HTML(200, "airdrop")
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

	if tu.Message.Chat.ID == -1001325718529 || tu.Message.Chat.ID == -1001161265502 || tu.Message.Chat.ID == -1001397587839 {
		lang = "hr"
	} else if tu.Message.Chat.ID == -1001249635625 {
		lang = "sr"
	} else {
		lang = "en-US"
	}

	log.Println(tu.Message.Chat.ID)

	if len(msgArr) == 1 {
		if msgArr[0] == "/gimme@AnonsRobot" {
			msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "addressRequired"))
		} else {
			send = false
		}
	} else if len(msgArr) > 1 {
		if msgArr[0] == "/gimme@AnonsRobot" {
			addr := msgArr[1]
			avr, err := wnc.AddressValidate(addr)
			if err == nil {
				if avr.Valid {
					user := &User{TelegramId: tu.Message.From.ID}
					db.First(user, user)

					if user.ID == 0 {
						if addr != conf.NodeAddress {
							user := &User{Address: addr}
							db.FirstOrCreate(user, user)
							if user.TelegramId == 0 {
								atr := &gowaves.AssetsTransferRequest{
									Amount:    100000000,
									AssetID:   "4zbprK67hsa732oSGLB6HzE8Yfdj3BcTcehCeTA1G5Lf",
									Fee:       100000,
									Recipient: addr,
									Sender:    conf.NodeAddress,
								}

								_, err := wnc.AssetsTransfer(atr)
								if err != nil {
									msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), fmt.Sprintf(ui18n.Tr(lang, "error"), err))
								} else {
									user.ReceivedFreeAnote = true
									user.TelegramId = tu.Message.From.ID
									db.Save(user)
									msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "anoteSent"))

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
							msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "yourAddress"))
						}
					} else if user.ReceivedFreeAnote {
						if user.Address != addr {
							msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "hacker"))
						} else {
							msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "alreadyActivated"))
						}
					}

					// user := &User{Address: addr}
					// db.FirstOrCreate(user, user)

					// if user.ID == 0

					// if user.ID == 0 {
					// 	user := &User{Address: addr}
					// 	db.FirstOrCreate(user, user)
					// 	if user.TelegramId == 0 {
					// 		user.TelegramId = tu.Message.From.ID
					// 		db.Save(user)
					// 	} else {
					// 		msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "hacker"))
					// 		bot.Send(msg)
					// 		return
					// 	}
					// } else if user.Address != addr {
					// 	msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "hacker"))
					// 	bot.Send(msg)
					// 	return
					// }

					// if user.ReceivedFreeAnote {
					// 	msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "alreadyActivated"))
					// } else {
					// 	if user.TelegramId == 0 {
					// 		atr := &gowaves.AssetsTransferRequest{
					// 			Amount:    100000000,
					// 			AssetID:   "4zbprK67hsa732oSGLB6HzE8Yfdj3BcTcehCeTA1G5Lf",
					// 			Fee:       100000,
					// 			Recipient: addr,
					// 			Sender:    conf.NodeAddress,
					// 		}

					// 		_, err := wnc.AssetsTransfer(atr)
					// 		if err != nil {
					// 			msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), fmt.Sprintf(ui18n.Tr(lang, "error"), err))
					// 		} else {
					// 			user.ReceivedFreeAnote = true
					// 			user.TelegramId = tu.Message.From.ID
					// 			db.Save(user)
					// 			msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "anoteSent"))

					// 			if len(user.Referral) > 0 {
					// 				atr := &gowaves.AssetsTransferRequest{
					// 					Amount:    20000000,
					// 					AssetID:   "4zbprK67hsa732oSGLB6HzE8Yfdj3BcTcehCeTA1G5Lf",
					// 					Fee:       100000,
					// 					Recipient: user.Referral,
					// 					Sender:    conf.NodeAddress,
					// 				}

					// 				wnc.AssetsTransfer(atr)
					// 			}
					// 		}
					// 	} else {
					// 		msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "hacker"))
					// 	}
					// }
				} else {
					msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), ui18n.Tr(lang, "addressNotValid"))
				}
			} else {
				msg = tgbotapi.NewMessage(int64(tu.Message.Chat.ID), fmt.Sprintf(ui18n.Tr(lang, "error"), err))
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

func view404(ctx *macaron.Context) {
	ctx.Data["Title"] = "Error 404 | "

	newPageData(ctx)

	ctx.HTML(200, "404")
}
