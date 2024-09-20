package vk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	methodGet  = "messages.getHistory"
	methodSend = "messages.send"
	vkBotHost  = "api.vk.com"
)

type Client struct {
	host   string
	token  string
	client http.Client
}

func New(token string) *Client {
	return &Client{
		host:   vkBotHost,
		token:  token,
		client: http.Client{},
	}
}

func (c *Client) Updates(offset, count int) ([]IncommingMessage, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("count", strconv.Itoa(count))
	q.Add("access_token", c.token)

	data, err := c.request(methodGet, q)
	if err != nil {
		return nil, err
	}

	var res []IncommingMessage

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) SendMessage(user_id, chat_id int, message string) error {
	q := url.Values{}
	q.Add("user_id", strconv.Itoa(user_id))
	q.Add("chat_id", strconv.Itoa(chat_id))
	q.Add("message", message)
	q.Add("access_token", c.token)

	_, err := c.request(methodSend, q)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) request(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join("method", method),
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Cant do request: %w", err)
	}

	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Cant get request-response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERR] Cant get response-body: %w", err)
	}
	return body, nil
}
