package main

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func newPageData(ctx *macaron.Context, sess session.Store) {
	ctx.Data["URI"] = ctx.Req.RequestURI
}
