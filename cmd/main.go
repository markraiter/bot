package main

import (
	"flag"
	"log"

	tgClient "github.com/markraiter/bot/pkg/clients/telegram"
	event_consumer "github.com/markraiter/bot/pkg/consumer/event_consumer"
	"github.com/markraiter/bot/pkg/events/telegram"
	"github.com/markraiter/bot/pkg/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, generateToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatalf("service is stopped: %s", err)
	}
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
