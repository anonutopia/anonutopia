package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anonutopia/gowaves"
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"gopkg.in/telegram-bot-api.v4"
)

var m *macaron.Macaron

var conf *Config

var bot *tgbotapi.BotAPI

var wnc *gowaves.WavesNodeClient

func main() {
	conf = initConfig()

	m = initMacaron()

	bot = initBot()

	wnc = initWaves()

	m.Get("/", newPageData, homeView)
	m.Get("/profit/", newPageData, profitView)
	m.Get("/anote/", newPageData, anoteView)
	m.Get("/projects/", newPageData, projectsView)
	m.Get("/contact/", newPageData, contactView)

	m.Post("/webhook/", newPageData, binding.Json(TelegramUpdate{}), webhookView)

	// m.NotFound(view404)

	// m.Run()
	log.Println("Server is running...")
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", 5000), m)
}
