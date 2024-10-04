package telegram

import (
	"BotStudyPSUTI/events"
	"BotStudyPSUTI/storage"
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
	textArray := strings.Split(text, ":")

	log.Printf("got new command '%s' from '%s", text, username)

	switch strings.TrimSpace(textArray[0]) {
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
		if len(textArray) == 2 {
			err := w.db.Save(textArray[1], &storage.UserInfo{Username: username, TypeApplication: storage.Tg})
			if err != nil {
				w.tg.SendMessage(chatID, events.MsgErrorOrder)
				return err
			}
			err = w.tg.SendMessage(chatID, events.MsgSaveOrder)
			return err
		}
		err := w.tg.SendMessage(chatID, events.MsgHelpOrder)
		return err
	default:
		err := w.tg.SendMessage(chatID, events.MsgUnknown)
		return err
	}
}
