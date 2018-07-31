package main

import (
	"strings"

	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func newPageData(ctx *macaron.Context, sess session.Store) {
	uri := strings.Split(ctx.Req.RequestURI, "?")
	ctx.Data["URI"] = uri[0]
}
