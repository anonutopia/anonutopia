package main

import (
	"gopkg.in/macaron.v1"
)

func homeView(ctx *macaron.Context) {
	ctx.Data["Title"] = ""

	ctx.HTML(200, "home")
}

func profitView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Making Profit in Anonutopia | "

	ctx.HTML(200, "profit")
}

func contactView(ctx *macaron.Context) {
	ctx.Data["Title"] = "Contact | "

	ctx.HTML(200, "contact")
}
