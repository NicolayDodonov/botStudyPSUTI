package vk

import (
	"BotStudyPSUTI/client/vk"
	"BotStudyPSUTI/events"
	"fmt"
)

type Worker struct {
	vk *vk.Client
}

type Meta struct {
	UserId int
}

func New(client *vk.Client) Worker {
	return Worker{
		vk: client,
	}
}

func (w *Worker) Fetch(limit int) ([]events.Event, error) {
	messages, err := w.vk.Updates()
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(messages))
	for _, m := range messages {
		if m.Out == 1 {
			continue
		}
		res = append(res, event(m))
	}
	return res, nil
}

func (w *Worker) Process(e events.Event) error {
	meta, err := meta(e)
	if err != nil {
		return fmt.Errorf("[ERR] cant worl meta %w", err)
	}

	if err := w.doCmd(e.Text, meta.UserId); err != nil {
		return fmt.Errorf("[ERR] cant understand command")
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

func event(m vk.Message) events.Event {
	res := events.Event{
		Type: events.Message,
		Text: m.Text,
		Meta: Meta{
			UserId: m.UserId,
		},
	}
	return res
}
