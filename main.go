package main

import (
	event_consumer "BotStudyPSUTI/Consumer/event-consumer"
	clientTg "BotStudyPSUTI/client/telegram"
	eventTg "BotStudyPSUTI/events/telegram"
	"flag"
	"log"
)

func main() {
	Event := eventTg.New(clientTg.New(mustToken()))

	log.Print("Service is started")
	Consumer := event_consumer.New(&Event, &Event, 100)
	if err := Consumer.Start(); err != nil {
		log.Fatal("Service is stopped")
	}
}

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"for access to bot api",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("[ERR] Token is not be found")
	}
	return *token
}
