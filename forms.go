package main

import (
	"github.com/go-macaron/binding"
	macaron "gopkg.in/macaron.v1"
)

type SignupForm struct {
	Email   string `form:"email" binding:"Required"`
	Seed    string `form:"seed" binding:"Required"`
	Address string `form:"address" binding:"Required"`
}

func (cf SignupForm) Error(ctx *macaron.Context, errs binding.Errors) {
	ctx.Data["Errors"] = errs
}
