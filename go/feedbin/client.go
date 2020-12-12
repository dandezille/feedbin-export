package feedbin

import (
  "log"
  "net/http"
  "io/ioutil"
  "strings"
  "encoding/json"
  "fmt"

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
  entries := getEntries(starred)
  return entries
}

func ensureAuthenticated() {
  request("authentication.json")
}

func getStarredEntries() []int {
  response := request("starred_entries.json")

  var starred []int
  err := json.Unmarshal(response, &starred)
  if err != nil {
    log.Fatal(err)
  }

  return starred
}

func getEntries(starred []int) []FeedEntry {
  ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(starred)), ","), "[]")
  response := request("entries.json?ids=" + ids)

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

func request(path string) []byte {
  url := url(path)
  log.Println("request: " + url)

  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal(err)
  }

  user := utils.ReadEnv("FEEDBIN_USER")
  password := utils.ReadEnv("FEEDBIN_PASSWORD")

  request.SetBasicAuth(user, password)
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

