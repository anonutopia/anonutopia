package main

import (
	"strings"

	"gopkg.in/macaron.v1"
)

func newPageData(ctx *macaron.Context) {
	uri := strings.Split(ctx.Req.RequestURI, "?")
	ctx.Data["URI"] = uri[0]
}
