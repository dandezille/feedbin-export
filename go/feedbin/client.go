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

  "github.com/dandezille/feedbin-to-todoist/utils"
)

type FeedbinClient struct {
}

type FeedEntry struct {
  Id  int `json:"id"`
  Url string `json:"url"`
}

func Connect() FeedbinClient {
  ensureAuthenticated()
  return FeedbinClient{}
}

func (c *FeedbinClient) GetStarredEntries() []FeedEntry {
  starred := getStarredEntries()
  if len(starred) == 0 {
    return nil
  }

  entries := getEntries(starred)
  return entries
}

func (c *FeedbinClient) Unstar(entries []FeedEntry) {
  var ids []string
  for _, entry := range entries {
    ids = append(ids, strconv.Itoa(entry.Id))
  }

  body := `{"starred_entries": [` + strings.Join(ids, ",") + "]}"
  log.Println(body)
  delete("starred_entries.json", body)
}

func ensureAuthenticated() {
  get("authentication.json")
}

func getStarredEntries() []int {
  response := get("starred_entries.json")

  var starred []int
  err := json.Unmarshal(response, &starred)
  if err != nil {
    log.Fatal(err)
  }

  return starred
}

func getEntries(starred []int) []FeedEntry {
  ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(starred)), ","), "[]")
  response := get("entries.json?ids=" + ids)

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

func get(path string) []byte {
  return request("GET", path, "")
}

func delete(path string, body string) []byte {
  return request("DELETE", path, body)
}

func request(method string, path string, data string) []byte {
  url := url(path)
  log.Println("request: " + url)

  request, err := http.NewRequest(method, url, getBody(data))
  if err != nil {
    log.Fatal(err)
  }

  user := utils.ReadEnv("FEEDBIN_USER")
  password := utils.ReadEnv("FEEDBIN_PASSWORD")

  request.SetBasicAuth(user, password)
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
