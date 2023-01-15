package main

import (
	"flag"
	"log"
	tgClient "tgGoLang/clients/telegram"
	"tgGoLang/events/telegram"
	"tgGoLang/consumer/eventConsumer"
	"tgGoLang/storage/files"
)

const (
	tgBotHost = "api.telegram.org"
	storagePath = "files_storage"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()), 
		files.New(storagePath),
	)

	log.Printf("service started")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}



	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token", 
		"", 
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not defined")
	}

	return *token
}