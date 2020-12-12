package main

import (
  "github.com/joho/godotenv"
  "encoding/json"
  "log"
  "strings"
  "fmt"

  "github.com/dandezille/feedbin-to-todoist/feedbin"
)

type StarredResult []int

type FeedEntry struct {
  Url   string `json:"url"`
}

type EntriesResult []FeedEntry

func ensureAuthenticated() {
  feedbin.Request("authentication.json")
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  ensureAuthenticated()
  response := feedbin.Request("starred_entries.json")

  var starred StarredResult
  err = json.Unmarshal(response, &starred)
  if err != nil {
    log.Fatal(err)
  }

  ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(starred)), ","), "[]")
  response = feedbin.Request("entries.json?ids=" + ids)

  var entries EntriesResult
  err = json.Unmarshal(response, &entries)
  if err != nil {
    log.Fatal(err)
  }

  log.Println(entries)
}
