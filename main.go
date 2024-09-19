package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	//todo: Токен Активации бота
	token := mustToken()
	fmt.Println(token)

	//todo: Активация Клиента API

	//todo: получатель данных
	//todo: обработчик данных
}

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"for acces to bot api",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("[ERR] Token is not be found")
	}
	return *token
}
