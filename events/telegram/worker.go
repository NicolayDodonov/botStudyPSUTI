package telegram

import (
	"BotStudyPSUTI/client/telegram"
	"BotStudyPSUTI/events"
	storage "BotStudyPSUTI/storage/sqlite"
	"fmt"
)

type Worker struct {
	tg     *telegram.Client
	db     *storage.SQLiteStorage
	offset int
}

type Meta struct {
	ChatId   int
	Username string
}

func New(client *telegram.Client, db *storage.SQLiteStorage) Worker {
	return Worker{
		tg: client,
		db: db,
	}
}

func (w *Worker) Fetch(limit int) ([]events.Event, error) {
	updates, err := w.tg.Updates(w.offset, limit)
	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))
	for _, u := range updates {
		res = append(res, event(u))
	}

	w.offset = updates[len(updates)-1].Id + 1
	return res, nil
}

func (w *Worker) Process(e events.Event) error {
	switch e.Type {
	case events.Message:
		if err := w.workMessage(e); err != nil {
			return err
		}
	default:
		return fmt.Errorf("[ERR] Cant work with message")
	}
	return nil
}

func (w *Worker) workMessage(e events.Event) error {
	meta, err := meta(e)
	if err != nil {
		return fmt.Errorf("[ERR] cant worl meta %w", err)
	}

	if err := w.doCmd(e.Text, meta.ChatId, meta.Username); err != nil {
		return fmt.Errorf("cant understand command")
	}

	return nil
}

func meta(e events.Event) (Meta, error) {
	res, ok := e.Meta.(Meta)
	if !ok {
		return Meta{}, fmt.Errorf("Cant take meta")
	}
	return res, nil
}

func event(u telegram.Update) events.Event {
	updType := fetchType(u)
	res := events.Event{
		Type: updType,
		Text: fetchText(u),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatId:   u.Message.Chat.Id,
			Username: u.Message.From.Username,
		}
	}

	return res
}

func fetchType(u telegram.Update) events.Type {
	if u.Message == nil {
		return events.Unknow
	}
	return events.Message
}

func fetchText(u telegram.Update) string {
	if u.Message == nil {
		return ""
	}
	return u.Message.Text
}
