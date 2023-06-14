package main

import (
	"flag"
	"log"

	"github.com/markraiter/bot/pkg/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, generateToken())

	// fetcher := fetcher.New(tgClient)

	// processor := processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func generateToken() string {
	token := flag.String(
		"telegram-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
