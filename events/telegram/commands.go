package telegram

import (
	"log"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (w *Worker) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	default:
		err := w.tg.SendMessage(chatID, "Я вас не понимаю")
		return err
	}
}
