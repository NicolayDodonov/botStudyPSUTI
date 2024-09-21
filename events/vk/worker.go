package vk

import (
	"BotStudyPSUTI/client/vk"
	"BotStudyPSUTI/events"
)

type Worker struct {
	vk *vk.Client
}

func New(client *vk.Client) Worker {
	return Worker{
		vk: client,
	}
}

func (w *Worker) Fetch(limit int) ([]events.Event, error) {
	return nil, nil
}

func (w *Worker) Process(e events.Event) error {
	return nil
}
