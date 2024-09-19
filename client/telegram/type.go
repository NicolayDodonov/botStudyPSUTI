package telegram

type UpdatesResponse struct {
	Ok     bool `json:"ok"`
	Result []Update
}

type Update struct {
	Id      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Username string `json:"first_name"`
}
type Chat struct {
	Id int `json:"id"`
}
