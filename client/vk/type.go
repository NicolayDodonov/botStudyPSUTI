package vk

type IncommingMessage struct {
	UserId int    `json:"from_id"`
	Text   string `json:"text"`
}
