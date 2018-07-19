package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/macaron.v1"
)

var m *macaron.Macaron

var conf *Config

func main() {
	conf = initConfig()

	m = initMacaron()

	m.Get("/", newPageData, homeView)

	// m.NotFound(view404)

	// m.Run()
	log.Println("Server is running...")
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", 5000), m)
}
