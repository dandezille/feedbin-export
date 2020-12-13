package feedbin

import (
  "log"
  "net/http"
  "io"
  "bytes"
  "io/ioutil"
  "strings"
  "encoding/json"
  "fmt"
  "strconv"
)

type Client struct {
  user string
  password string
}

type FeedEntry struct {
  Id  int `json:"id"`
  Url string `json:"url"`
}

func Connect(user string, password string) Client {
  c := Client{
    user: user,
    password: password,
  }
  c.ensureAuthenticated()
  return c
}

func (c *Client) GetStarredEntries() []FeedEntry {
  starred := c.getStarredEntries()
  if len(starred) == 0 {
    return nil
  }

  entries := c.getEntries(starred)
  return entries
}

func (c *Client) Unstar(entries []FeedEntry) {
  var ids []string
  for _, entry := range entries {
    ids = append(ids, strconv.Itoa(entry.Id))
  }

  body := `{"starred_entries": [` + strings.Join(ids, ",") + "]}"
  log.Println(body)
  c.delete("starred_entries.json", body)
}

func (c *Client) ensureAuthenticated() {
  c.get("authentication.json")
}

func (c *Client) getStarredEntries() []int {
  response := c.get("starred_entries.json")

  var starred []int
  err := json.Unmarshal(response, &starred)
  if err != nil {
    log.Fatal(err)
  }

  return starred
}

func (c *Client) getEntries(starred []int) []FeedEntry {
  ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(starred)), ","), "[]")
  response := c.get("entries.json?ids=" + ids)

  var entries []FeedEntry
  err := json.Unmarshal(response, &entries)
  if err != nil {
    log.Fatal(err)
  }

  return entries
}

func url(path string) string {
  baseUrl := "https://api.feedbin.com/v2/"
  return baseUrl + path
}

func (c *Client) get(path string) []byte {
  return c.request("GET", path, "")
}

func (c *Client) delete(path string, body string) []byte {
  return c.request("DELETE", path, body)
}

func (c *Client) request(method string, path string, data string) []byte {
  url := url(path)
  log.Println("request: " + url)

  request, err := http.NewRequest(method, url, getBody(data))
  if err != nil {
    log.Fatal(err)
  }

  request.SetBasicAuth(c.user, c.password)
  request.Header.Add("Content-Type", "application/json")

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

func getBody(data string) io.Reader {
  if data == "" {
    return nil
  } else {
    return bytes.NewBuffer([]byte(data))
  }
}
