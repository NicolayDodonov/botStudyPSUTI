package vk

type LongPollyResponse struct {
	Ts int `json:"ts"`
}

type LongPollyUpdate struct {
	MessageArray MessageArray `json:"messages"`
}

type MessageArray struct {
	Messages []Message `json:"items"`
}

type Message struct {
	UserId int    `json:"from_id"`
	Out    int    `json:"out"`
	Text   string `json:"text"`
}
