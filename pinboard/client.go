package pinboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dandezille/feedbin-export/utils"
)

type Client struct {
	client utils.RestClient
}

type Result struct {
	Code string `json:"result_code"`
}

func Connect(key string) (*Client, error) {
	customiseRequest := func(request *http.Request) {
		q := request.URL.Query()
		q.Add("auth_token", key)
		q.Add("format", "json")
		request.URL.RawQuery = q.Encode()
	}

	c := &Client{
		client: utils.NewRestClient("https://api.pinboard.in/v1/", customiseRequest),
	}

	err := c.authenticate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) CreateEntry(content string, title string) error {
	q := fmt.Sprintf("posts/add?url=%v&description=%v&toread=yes", url.QueryEscape(content), url.QueryEscape(title))

	response, err := c.client.Get(q)
	if err != nil {
		return err
	}

	var result Result
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.Code != "done" {
		return fmt.Errorf("Unexpected result code: \"%v\"", result.Code)
	}

	return nil
}

func (c *Client) authenticate() error {
	_, err := c.client.Get("user/api_token")
	return err
}
