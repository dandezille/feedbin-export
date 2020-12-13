package todoist

import (
  "log"
  "net/http"
  "io/ioutil"

  "github.com/dandezille/feedbin-to-todoist/utils"
)

type Client struct {
  key string
  client utils.RestClient
}

func Connect(key string) Client {
  c := Client{
    key: key,
    client: utils.NewRestClient("https://api.todoist.com/rest/v1/"),
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

func (c *Client) newRequest(method string, path string, data string) *http.Request {
  request, err := http.NewRequest(method, path, utils.BodyFromString(data))
  if err != nil {
    log.Fatal(err)
  }

  request.Header.Add("Authorization", "Bearer " + c.key)
  request.Header.Add("Content-Type", "application/json")

  return request
}

func (c *Client) request(method string, path string, data string) []byte {
  url := c.client.Url(path)
  log.Println("request: " + url)

  request := c.newRequest(method, url, data)

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  if response.StatusCode != 200 {
    log.Fatal("request for " + url + " status " + response.Status)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }

  log.Println("response: " + string(body))
  return body
}
