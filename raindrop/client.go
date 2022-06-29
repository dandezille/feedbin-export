package raindrop

import (
	"encoding/json"
	"net/http"

	"github.com/dandezille/feedbin-export/utils"
)

type Client struct {
	client utils.RestClient
}

type Request struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type Result struct {
	Ok bool `json:"result"`
}

func Connect(key string) (*Client, error) {
	customiseRequest := func(request *http.Request) {
		request.Header.Add("Authorization", "Bearer "+key)
	}

	c := &Client{
		client: utils.NewRestClient("https://api.raindrop.io/rest/v1", customiseRequest),
	}

	err := c.authenticate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) CreateEntry(content string, title string) error {
	data := Request{
		Title: title,
		Link:  content,
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = c.client.Post("/raindrop", string(dataJSON))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) authenticate() error {
	_, err := c.client.Get("/user")
	return err
}
