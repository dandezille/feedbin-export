package main

import (
  "github.com/joho/godotenv"
  "io/ioutil"
  "encoding/json"
  "log"
  "net/http"
  "os"
  "strings"
  "fmt"
)

type StarredResult []int

type FeedEntry struct {
  Url   string `json:"url"`
}

type EntriesResult []FeedEntry


func readEnv(name string) string {
  value, present := os.LookupEnv(name)
  if !present {
    log.Fatal(name + "not present")
  }

  return value
}

func feedbinUrl(path string) string {
  baseUrl := "https://api.feedbin.com/v2/"
  return baseUrl + path
}

func feedbinRequest(path string) []byte {
  url := feedbinUrl(path)
  log.Println("request: " + url)

  request, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal(err)
  }

  request.SetBasicAuth(readEnv("FEEDBIN_USER"), readEnv("FEEDBIN_PASSWORD"))
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

func ensureAuthenticated() {
  feedbinRequest("authentication.json")
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  ensureAuthenticated()
  response := feedbinRequest("starred_entries.json")

  var starred StarredResult
  err = json.Unmarshal(response, &starred)
  if err != nil {
    log.Fatal(err)
  }

  ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(starred)), ","), "[]")
  response = feedbinRequest("entries.json?ids=" + ids)

  var entries EntriesResult
  err = json.Unmarshal(response, &entries)
  if err != nil {
    log.Fatal(err)
  }

  log.Println(entries)
}
