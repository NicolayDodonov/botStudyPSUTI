package telegram

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
	methodGet  = "getUpdates"
	methodSend = "sendMessage"
	tgBotHost  = "api.telegram.org"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(token string) *Client {
	return &Client{
		host:     tgBotHost,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

// Метод для получения сообщений
func (c *Client) Updates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.request(methodGet, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

// Метод для отправки сообщений
func (c *Client) SendMessage(chatId int, text string) error {
	q := url.Values{}
	q.Add("chatId", strconv.Itoa(chatId))
	q.Add("text", text)

	_, err := c.request(methodSend, q)
	if err != nil {
		return err
	}

	return nil
}

// Метод для запросов API
func (c *Client) request(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
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

func newBasePath(token string) string {
	return "bot" + token
}
