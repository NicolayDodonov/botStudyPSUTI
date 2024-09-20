package telegram

import (
	"BotStudyPSUTI/events"
	"log"
	"strings"
)

const (
	CmdStart = "/start"
	CmdMenu  = "/menu"
	CmdHelp  = "/help"
	CmdOrder = "/order"
)

func (w *Worker) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	case CmdStart:
		err := w.tg.SendMessage(chatID, events.MsgStart)
		return err
	case CmdMenu:
		err := w.tg.SendMessage(chatID, events.MsgMenu)
		return err
	case CmdHelp:
		err := w.tg.SendMessage(chatID, events.MsgHelp)
		return err
	case CmdOrder:
		err := w.tg.SendMessage(chatID, events.MsgOrder)
		return err
	default:
		err := w.tg.SendMessage(chatID, events.MsgUnknown)
		return err
	}
}
