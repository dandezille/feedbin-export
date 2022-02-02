package todoist

import (
	"net/http"

	"github.com/dandezille/feedbin-export/utils"
)

type Client struct {
	key    string
	client utils.RestClient
}

func Connect(key string) Client {
	customiseRequest := func(request *http.Request) {
		request.Header.Add("Authorization", "Bearer "+key)
	}

	c := Client{
		key:    key,
		client: utils.NewRestClient("https://api.todoist.com/rest/v1/", customiseRequest),
	}

	c.ensureAuthenticated()
	return c
}

func (c *Client) CreateEntry(content string) {
	c.client.Post("tasks", `{ "content": "`+content+`" }`)
}

func (c *Client) ensureAuthenticated() {
	c.client.Get("projects")
}
