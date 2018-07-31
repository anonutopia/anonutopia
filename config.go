package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Redis           string `json:"redis"`
	Port            uint   `json:"port"`
	Telegram        string `json:"telegram"`
	WavesNodeApiKey string `json:"wavesnode_apikey"`
	Debug           bool   `json:"debug"`
	NodeAddress     string `json:"node_address"`
}

func (sc *Config) Load(configFile string) error {
	file, err := os.Open(configFile)

	if err != nil {
		log.Printf("[Config.Load] Got error while opening config file: %v", err)
		return err
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&sc)

	if err != nil {
		log.Printf("[Config.Load] Got error while decoding JSON: %v", err)
		return err
	}

	return nil
}

func initConfig() *Config {
	c := &Config{}
	c.Load("config.json")
	return c
}
