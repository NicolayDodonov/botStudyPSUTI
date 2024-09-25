package vk

type LongPollyConnect struct {
	Response ResponseConnect `json:"response"`
	Error    Error           `json:"error"`
}

type ResponseConnect struct {
	Server string `json:"server"`
	Key    string `json:"key"`
	Ts     int    `json:"ts"`
}

type Error struct {
	Code int    `json:"error_code"`
	Text string `json:"error_msg"`
}

//*==========================*//

type LongPollyUpdate struct {
	Response ResponseUpdate `json:"response"`
	Error    Error          `json:"error"`
}

type ResponseUpdate struct {
	MessageArray MessageArray `json:"messages"`
}

type MessageArray struct {
	Count    int       `json:"count"`
	Messages []Message `json:"items"`
}

type Message struct {
	UserId int    `json:"from_id"`
	Out    int    `json:"out"`
	Text   string `json:"text"`
}
