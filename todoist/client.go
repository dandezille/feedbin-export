package todoist

import (
	"log"
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
	_, err := c.client.Post("tasks", `{ "content": "`+content+`" }`)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Client) ensureAuthenticated() {
	_, err := c.client.Get("projects")
	if err != nil {
		log.Fatal(err)
	}
}
