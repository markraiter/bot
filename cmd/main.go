package main

import (
	"context"
	"flag"
	"log"

	_ "github.com/mattn/go-sqlite3"

	tgClient "github.com/markraiter/bot/pkg/clients/telegram"
	event_consumer "github.com/markraiter/bot/pkg/consumer/event_consumer"
	"github.com/markraiter/bot/pkg/events/telegram"
	"github.com/markraiter/bot/pkg/storage/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	sqlitestoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	// s := files.New(storagePath)
	s, err := sqlite.New(sqlitestoragePath)
	if err != nil {
		log.Fatalf("can't connect to storage: %s", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatalf("can't init storage: %s", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, generateToken()),
		s,
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
