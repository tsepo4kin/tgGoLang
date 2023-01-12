package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	storagePath = "/telegram"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()), 
		files.New(storagePath)
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
		"token-bot-token", 
		"", 
		"token for access to telegram bot"
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not defined")
	}

	return token
}