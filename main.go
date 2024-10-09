package main

import (
	eventConsumer "BotStudyPSUTI/consumer/event-consumer"
	clientTg "BotStudyPSUTI/client/telegram"
	clientVk "BotStudyPSUTI/client/vk"
	eventTg "BotStudyPSUTI/events/telegram"
	eventVk "BotStudyPSUTI/events/vk"
	storageSQLite "BotStudyPSUTI/storage/sqlite"
	"flag"
	"log"
)

func main() {

	token, types := mustFlag()

	storage, err := storageSQLite.New("storage/data/database.db")
	if err != nil {
		log.Fatal(err)
	}

	err = storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	switch types {
	case "tg":
		Worker := eventTg.New(clientTg.New(token), storage)
		log.Print("Telegram service is started")

		Consumer := eventConsumer.New(&Worker, &Worker, 100)
		if err := Consumer.Start(); err != nil {
			log.Fatal("Telegram service is stopped")
		}
	case "vk":
		Worker := eventVk.New(clientVk.New(token))
		log.Print("Vk service is started")

		Consumer := eventConsumer.New(&Worker, &Worker, 100)
		if err := Consumer.Start(); err != nil {
			log.Fatal("Telegram service is stopped")
		}
	default:
		log.Fatal("[ERR] Сan't determine the type app")
	}

}

func mustFlag() (string, string) {
	token := flag.String(
		"token",
		"",
		"for access to bot api",
	)

	types := flag.String(
		"types",
		"",
		"for select type apps",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("[ERR] Token is not be found")
	}
	if *types == "" {
		log.Fatal("[ERR] Types is not be found")
	}
	return *token, *types
}
