package main

import (
	"fmt"
	"time"
)

type Anon struct {
	Applicants string
	Citizens   string
	Founders   string
	Budget     string
}

func (a *Anon) loadData() {
	applicant := &Badge{Name: "applicant"}
	db.Preload("Users").First(applicant, applicant)

	citizen := &Badge{Name: "citizen"}
	db.Preload("Users").First(citizen, citizen)

	founder := &Badge{Name: "founder"}
	db.Preload("Users").First(founder, founder)

	prices, err := pc.DoRequest()
	anote.loadState()
	if err == nil {
		amount := uint64(0)

		amount += anote.BudgetWav / uint64(prices.WAVES*float64(satInBtc))
		amount += anote.BudgetBtc / uint64(prices.BTC*float64(satInBtc))
		amount += anote.BudgetEth / uint64(prices.ETH*float64(satInBtc))

		a.Budget = fmt.Sprintf("%d", amount)
	} else {
		a.Budget = "0.00"
	}

	a.Applicants = fmt.Sprintf("%d", len(applicant.Users))
	a.Citizens = fmt.Sprintf("%d", len(citizen.Users))
	a.Founders = fmt.Sprintf("%d", len(founder.Users))
}

func (a *Anon) sendEmails() {
	var users []*User
	db.Find(&users)
	for _, u := range users {
		if u.EmailVerified && u.CreatedAt.Add(time.Hour*24).Before(time.Now()) && u.HasBadges() && !u.SentFollowUp {
			sendFollowUpEmail(u, "en-US")
			u.SentFollowUp = true
			db.Save(u)
		}
	}
}

func initAnon() *Anon {
	anon := &Anon{}

	go func() {
		for {
			anon.loadData()
			anon.sendEmails()
			time.Sleep(5 * time.Minute)
		}
	}()

	return anon
}
