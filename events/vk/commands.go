package vk

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

func (w *Worker) doCmd(text string, userId int) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%v", text, userId)

	switch text {
	case CmdStart:
		err := w.vk.SendMessage(userId, events.MsgStart)
		return err
	case CmdMenu:
		err := w.vk.SendMessage(userId, events.MsgMenu)
		return err
	case CmdHelp:
		err := w.vk.SendMessage(userId, events.MsgHelp)
		return err
	case CmdOrder:
		err := w.vk.SendMessage(userId, events.MsgOrder)
		return err
	default:
		err := w.vk.SendMessage(userId, events.MsgUnknown)
		return err
	}
}
