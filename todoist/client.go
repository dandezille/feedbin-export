package todoist

import (
  "net/http"

  "github.com/dandezille/feedbin-to-todoist/utils"
)

type Client struct {
  key string
  client utils.RestClient
}

func Connect(key string) Client {
  customiseRequest := func(request *http.Request) {
    request.Header.Add("Authorization", "Bearer " + key)
  }

  c := Client{
    key: key,
    client: utils.NewRestClient("https://api.todoist.com/rest/v1/", customiseRequest),
  }

  c.ensureAuthenticated()
  return c
}

func (c *Client) CreateEntry(content string) {
  c.post("tasks", `{ "content": "` + content + `" }`)
}

func (c *Client) ensureAuthenticated() {
  c.get("projects")
}

func (c *Client) get(path string) {
  c.request("GET", path, "")
}

func (c *Client) post(path string, data string) {
  c.request("POST", path, data)
}

func (c *Client) request(method string, path string, data string) []byte {
  request := c.client.NewRequest(method, path, data)
  return c.client.Execute(request)
}
