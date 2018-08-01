package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Nickname                string `sql:"size:255"`
	Email                   string `sql:"size:255"`
	Address                 string `sql:"size:255;unique_index"`
	Referral                string `sql:"size:255"`
	BitcoinAddr             string `sql:"size:255"`
	BitcoinBalanceNew       int    `sql:"DEFAULT:0"`
	BitcoinBalanceProcessed int    `sql:"DEFAULT:0"`
	EtherAddr               string `sql:"size:255"`
	EtherBalanceNew         int    `sql:"DEFAULT:0"`
	EtherBalanceProcessed   int    `sql:"DEFAULT:0"`
	ProfitEth               uint64
	ProfitWav               uint64
	ProfitBtc               uint64
	ProfitEthTotal          uint64
	ProfitWavTotal          uint64
	ProfitBtcTotal          uint64
	ReferralProfitEth       uint64
	ReferralProfitWav       uint64
	ReferralProfitBtc       uint64
	ReferralProfitEthTotal  uint64
	ReferralProfitWavTotal  uint64
	ReferralProfitBtcTotal  uint64
	ReceivedFreeAnote       bool `sql:"DEFAULT:false"`
	TelegramId              int  `sql:"DEFAULT:0;unique_index"`
}
