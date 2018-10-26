package main

import (
	"fmt"
	"strings"

	"gopkg.in/macaron.v1"
)

func newPageData(ctx *macaron.Context) {
	anote.loadState()
	uri := strings.Split(ctx.Req.RequestURI, "?")
	ctx.Data["URI"] = uri[0]
	ctx.Data["NodeAddress"] = conf.NodeAddress
	ctx.Data["Anon"] = anon
	ctx.Data["PriceEur"] = fmt.Sprintf("%.8f", float64(anote.Price)/float64(satInBtc))
}
