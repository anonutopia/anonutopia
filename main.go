package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anonutopia/gowaves"
	"github.com/go-macaron/binding"
	"github.com/jinzhu/gorm"
	macaron "gopkg.in/macaron.v1"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var m *macaron.Macaron

var conf *Config

var bot *tgbotapi.BotAPI

var wnc *gowaves.WavesNodeClient

var db *gorm.DB

var anon *Anon

var anote *Anote

var pc *PriceClient

func main() {
	conf = initConfig()

	m = initMacaron()

	bot = initBot()

	wnc = initWaves()

	db = initDb()

	anote = initAnote()

	anon = initAnon()

	pc = initPriceClient()

	m.Get("/", newPageData, homeView)
	m.Get("/crowdfunding/", newPageData, crowdfundingView)
	m.Get("/about/", newPageData, aboutView)
	m.Get("/faq/", newPageData, faqView)
	m.Get("/apply-advanced/", newPageData, applyAdvancedView)
	m.Get("/cryptocountries/", newPageData, cryptocountriesView)
	m.Get("/tags/", newPageData, tagsView)
	m.Get("/docs/", newPageData, docsView)
	m.Get("/staking/", newPageData, stakingView)
	m.Get("/airdrop/", newPageData, airdropView)
	m.Get("/profit/", newPageData, profitView)
	m.Get("/anote/", newPageData, anoteView)
	m.Get("/projects/", newPageData, projectsView)
	m.Get("/contact/", newPageData, contactView)
	m.Get("/pressrelease/", newPageData, prView)

	m.Post("/webhook/", newPageData, binding.Json(TelegramUpdate{}), webhookView)

	m.NotFound(view404)

	// m.Run()
	log.Println("Server is running...")
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", 5000), m)
}
