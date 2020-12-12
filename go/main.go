package main

import (
  "github.com/joho/godotenv"
  "log"

  "github.com/dandezille/feedbin-to-todoist/feedbin"
  "github.com/dandezille/feedbin-to-todoist/todoist"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  fetchFeeds()
}

func fetchFeeds() {
  feedbin := feedbin.Connect()
  entries := feedbin.GetStarredEntries()
  log.Println(entries)

  if len(entries) == 0 {
    log.Println("No starred entries")
    return
  }

  todoist := todoist.Connect()
  for _, entry := range entries {
    todoist.CreateEntry(entry.Url)
  }

  feedbin.Unstar(entries)
}
